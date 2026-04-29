package repositories

import (
	"context"
	"errors"

	"github.com/TaisukeFujise/blog_api/models"
)

type ArticleRepository interface {
	InsertArticle(ctx context.Context, article models.Article) (models.Article, error)
	SelectArticleList(ctx context.Context, page int) ([]models.Article, error)
	SelectArticleDetail(ctx context.Context, articleID int) (models.Article, error)
	UpdateNiceNum(ctx context.Context, articleID int) error
}

type CommentRepository interface {
	InsertComment(ctx context.Context, comment models.Comment) (models.Comment, error)
	SelectCommentList(ctx context.Context, articleID int) ([]models.Comment, error)
}

var ErrNotFound = errors.New("record not found")
