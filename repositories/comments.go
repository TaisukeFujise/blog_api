package repositories

import (
	"context"
	"database/sql"

	"github.com/TaisukeFujise/blog_api/models"
)

type CommentRepositoryImpl struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{db: db}
}

func (r *CommentRepositoryImpl) InsertComment(ctx context.Context, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Comment{}, err
	}

	_, err = tx.ExecContext(ctx, sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		tx.Rollback()
		return models.Comment{}, err
	}

	tx.Commit()
	return comment, nil
}

func (r *CommentRepositoryImpl) SelectCommentList(ctx context.Context, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`
	commentArray := make([]models.Comment, 0)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, sqlStr, articleID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime

		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if err != nil {
			tx.Rollback()
			return nil, err
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
