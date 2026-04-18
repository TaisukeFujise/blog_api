package services

import (
	"github.com/TaisukeFujise/blog_api/models"
	"github.com/TaisukeFujise/blog_api/repositories"
)

// -> ArticleDetailHandler
// 指定IDの記事情報を返却
func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()
	// 1. repositories.SelectArticleDetailで記事の詳細を取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 2. repositories.SelectCommentListでコメント一覧を取得
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 3. 2で得たコメント一覧を、1で得たArticle構造体に紐づける。
	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

// -> PostArticleHandler
// 引数のarticleを元に新しい記事を作成し、結果を返却
func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

// -> ArticleListHandler
// 指定pageの記事一覧を返却
func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return []models.Article{}, err
	}
	defer db.Close()

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return []models.Article{}, err
	}
	return articleList, nil
}

// -> PostNiceHandler
// 指定IDの記事のいいね数を+1して、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}
	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
