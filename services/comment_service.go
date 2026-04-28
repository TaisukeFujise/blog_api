package services

import (
	"context"

	"github.com/TaisukeFujise/blog_api/apperrors"
	"github.com/TaisukeFujise/blog_api/models"
)

func (s *CommentService) PostCommentService(ctx context.Context, comment models.Comment) (models.Comment, error) {
	comment, err := s.commentRepo.InsertComment(ctx, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}
	return comment, nil
}
