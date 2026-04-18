package services

import "github.com/TaisukeFujise/blog_api/models"

// type MyAppServicer interface {
// 	PostArticleService(article models.Article) (models.Article, error)
// 	GetArticleListService(page int) ([]models.Article, error)
// 	GetArticleService(articleID int) (models.Article, error)
// 	PostNiceService(article models.Article) (models.Article, error)
// 	PostCommentService(comment models.Comment) (models.Comment, error)
// }

// article関連を受けるservice
type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

// comment関連を受けるservice
type CommentServicer interface {
	PostCommentService(commnet models.Comment) (models.Comment, error)
}
