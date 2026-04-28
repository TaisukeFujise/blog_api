package services

import (
	"github.com/TaisukeFujise/blog_api/services/repositories"
)

type ArticleService struct {
	articleRepo repositories.ArticleRepository
	commentRepo repositories.CommentRepository
}

type CommentService struct {
	commentRepo repositories.CommentRepository
}

func NewArticleService(articleRepo repositories.ArticleRepository, commentRepo repositories.CommentRepository) *ArticleService {
	return &ArticleService{articleRepo: articleRepo, commentRepo: commentRepo}
}

func NewCommentService(commentRepo repositories.CommentRepository) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}
