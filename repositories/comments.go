package repositories

import (
	"database/sql"

	"github.com/TaisukeFujise/blog_api/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`

	tx, err := db.Begin()
	if err != nil {
		return comment, err
	}

	_, err = tx.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		tx.Rollback()
		return comment, err
	}

	tx.Commit()
	return comment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`
	commentArray := make([]models.Comment, 0)

	tx, err := db.Begin()
	if err != nil {
		return commentArray, err
	}

	rows, err := tx.Query(sqlStr, articleID)
	if err != nil {
		tx.Rollback()
		return commentArray, err
	}

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime

		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if err != nil {
		} else {
			if createdTime.Valid {
				comment.CreatedAt = createdTime.Time
			}
			commentArray = append(commentArray, comment)
		}
	}

	tx.Commit()
	return commentArray, nil
}
