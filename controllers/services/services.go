package services

import (
	"context"

	"github.com/TaisukeFujise/blog_api/models"
)

// article関連を受けるservice
type ArticleServicer interface {
	PostArticleService(ctx context.Context, article models.Article) (models.Article, error)
	GetArticleListService(ctx context.Context, page int) ([]models.Article, error)
	GetArticleService(ctx context.Context, articleID int) (models.Article, error)
	PostNiceService(ctx context.Context, article models.Article) (models.Article, error)
}

// comment関連を受けるservice
type CommentServicer interface {
	PostCommentService(ctx context.Context, comment models.Comment) (models.Comment, error)
}
