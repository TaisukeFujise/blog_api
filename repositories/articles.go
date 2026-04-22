package repositories

import (
	"context"
	"database/sql"

	"github.com/TaisukeFujise/blog_api/models"
)

func InsertArticle(ctx context.Context, db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return article, err
	}
	result, err := tx.ExecContext(ctx, sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		tx.Rollback()
		return article, err
	}

	tx.Commit()

	id, err := result.LastInsertId()
	if err != nil {
		return article, err
	}
	article.ID = int(id)
	return article, nil
}

func SelectArticleList(ctx context.Context, db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	rows, err := tx.QueryContext(ctx, sqlStr, 5, (page-1)*5)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article

		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
		} else {
			articleArray = append(articleArray, article)
		}
	}

	tx.Commit()
	return articleArray, nil
}

func SelectArticleDetail(ctx context.Context, db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	var article models.Article
	var createdTime sql.NullTime

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return article, err
	}

	row := tx.QueryRowContext(ctx, sqlStr, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return article, err
	}

	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	if err != nil {
		tx.Rollback()
		return article, err
	}

	tx.Commit()
	return article, nil
}

func UpdateNiceNum(ctx context.Context, db *sql.DB, articleID int) error {
	const sqlUpdateNice = `update articles set nice = nice + 1 where article_id = ?`

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, sqlUpdateNice, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
