package services

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/TaisukeFujise/blog_api/apperrors"
	"github.com/TaisukeFujise/blog_api/models"
	"github.com/TaisukeFujise/blog_api/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	// Lock周り
	var amu sync.Mutex
	var cmu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		amu.Lock()
		article, articleGetErr = repositories.SelectArticleDetail(db, articleID)
		amu.Unlock()
	}(s.db, articleID)
	// article, err := repositories.SelectArticleDetail(s.db, articleID)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		err = apperrors.NAData.Wrap(err, "no data")
	// 		return models.Article{}, err
	// 	}
	// 	err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
	// 	return models.Article{}, err
	// }

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		cmu.Lock()
		commentList, commentGetErr = repositories.SelectCommentList(db, articleID)
		cmu.Unlock()
	}(s.db, articleID)
	// commentList, err := repositories.SelectCommentList(s.db, articleID)
	// if err != nil {
	// 	err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
	// 	return models.Article{}, err
	// }

	wg.Wait()

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}
	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return []models.Article{}, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
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
