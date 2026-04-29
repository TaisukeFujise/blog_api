package repositories

import (
	"context"
	"database/sql"

	"github.com/TaisukeFujise/blog_api/models"
)

type ArticleRepositoryImpl struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepositoryImpl {
	return &ArticleRepositoryImpl{db: db}
}

func (r *ArticleRepositoryImpl) InsertArticle(ctx context.Context, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Article{}, err
	}
	result, err := tx.ExecContext(ctx, sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}

	tx.Commit()

	id, err := result.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}
	article.ID = int(id)
	return article, nil
}

func (r *ArticleRepositoryImpl) SelectArticleList(ctx context.Context, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`

	tx, err := r.db.BeginTx(ctx, nil)
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
			tx.Rollback()
			return nil, err
		} else {
			articleArray = append(articleArray, article)
		}
	}

	tx.Commit()
	return articleArray, nil
}

func (r *ArticleRepositoryImpl) SelectArticleDetail(ctx context.Context, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	var article models.Article
	var createdTime sql.NullTime

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Article{}, err
	}

	row := tx.QueryRowContext(ctx, sqlStr, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return models.Article{}, err
	}

	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}

	tx.Commit()
	return article, nil
}

func (r *ArticleRepositoryImpl) UpdateNiceNum(ctx context.Context, articleID int) error {
	const sqlUpdateNice = `update articles set nice = nice + 1 where article_id = ?`

	tx, err := r.db.BeginTx(ctx, nil)
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
