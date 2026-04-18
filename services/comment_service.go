package services

import (
	"github.com/TaisukeFujise/blog_api/models"
	"github.com/TaisukeFujise/blog_api/repositories"
)

// -> PostCommentHandler
// 引数のcommentを元に新しいコメントを作成し、結果を返却
func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	comment, err = repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}
