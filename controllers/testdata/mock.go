package testdata

import (
	"context"

	"github.com/TaisukeFujise/blog_api/models"
)

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostArticleService(ctx context.Context, article models.Article) (models.Article, error) {
	return articleTestData[1], nil
}

func (s *serviceMock) GetArticleListService(ctx context.Context, page int) ([]models.Article, error) {
	return articleTestData, nil
}

func (s *serviceMock) GetArticleService(ctx context.Context, articleID int) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) PostNiceService(ctx context.Context, article models.Article) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) PostCommentService(ctx context.Context, comment models.Comment) (models.Comment, error) {
	return commentTestData[0], nil
}
