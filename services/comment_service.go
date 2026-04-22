package services

import (
	"context"

	"github.com/TaisukeFujise/blog_api/apperrors"
	"github.com/TaisukeFujise/blog_api/models"
	"github.com/TaisukeFujise/blog_api/repositories"
)

func (s *MyAppService) PostCommentService(ctx context.Context, comment models.Comment) (models.Comment, error) {
	comment, err := repositories.InsertComment(ctx, s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}
	return comment, nil
}
