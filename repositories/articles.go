package repositories

import (
	"database/sql"

	"github.com/TaisukeFujise/blog_api/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`

	tx, err := db.Begin()
	if err != nil {
		return article, err
	}
	result, err := tx.Exec(sqlStr, article.Title, article.Contents, article.UserName)
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

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query(sqlStr, 5, (page-1)*5)
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

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	var article models.Article
	var createdTime sql.NullTime

	tx, err := db.Begin()
	if err != nil {
		return article, err
	}

	row := tx.QueryRow(sqlStr, articleID)
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

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
