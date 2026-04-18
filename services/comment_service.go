package services

import (
	"github.com/TaisukeFujise/blog_api/models"
	"github.com/TaisukeFujise/blog_api/repositories"
)

// -> PostCommentHandler
// 引数のcommentを元に新しいコメントを作成し、結果を返却
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	comment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}
