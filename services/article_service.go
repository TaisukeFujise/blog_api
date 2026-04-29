package services

import (
	"context"
	"errors"

	"github.com/TaisukeFujise/blog_api/apperrors"
	"github.com/TaisukeFujise/blog_api/models"
	"github.com/TaisukeFujise/blog_api/services/repositories"
)

func (s *ArticleService) GetArticleService(ctx context.Context, articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	type articleResult struct {
		article models.Article
		err     error
	}
	articleChan := make(chan articleResult)
	defer close(articleChan)
	go func(ch chan<- articleResult, articleID int) {
		article, err := s.articleRepo.SelectArticleDetail(ctx, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan, articleID)

	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}
	commentChan := make(chan commentResult)
	defer close(commentChan)
	go func(ch chan<- commentResult, articleID int) {
		commentList, err := s.commentRepo.SelectCommentList(ctx, articleID)
		ch <- commentResult{commentList: &commentList, err: err}
	}(commentChan, articleID)

	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		case <-ctx.Done():
			return models.Article{}, ctx.Err()
		}
	}

	if articleGetErr != nil {
		if errors.Is(articleGetErr, repositories.ErrNotFound) {
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

func (s *ArticleService) PostArticleService(ctx context.Context, article models.Article) (models.Article, error) {
	newArticle, err := s.articleRepo.InsertArticle(ctx, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *ArticleService) GetArticleListService(ctx context.Context, page int) ([]models.Article, error) {
	articleList, err := s.articleRepo.SelectArticleList(ctx, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return articleList, nil
}

func (s *ArticleService) PostNiceService(ctx context.Context, article models.Article) (models.Article, error) {
	err := s.articleRepo.UpdateNiceNum(ctx, article.ID)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
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
