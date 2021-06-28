package usecase

import (
	"fmt"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/logger"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/app/errors"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/services"
	"github.com/sirupsen/logrus"
)

type articleUsecase struct{}

func NewArticleUsecase() *articleUsecase {
	return &articleUsecase{}
}

func (u *articleUsecase) Get(ctx services.ArticleServiceContext, id uint) (*domain.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "get")
	logOp.Infof("Request to get an article: %d", id)
	result, err := ctx.ArticleRepository().FindByID(ctx, id)
	if err != nil {
		errorMsg := fmt.Sprintf("error getting an article with id %d", id)
		return nil, errors.NewArticleError(errors.ArticleRetrievalErrorCode, errorMsg, err)
	}
	return result, nil
}

func (u *articleUsecase) Find(ctx services.ArticleServiceContext, query *domain.ArticleQuery) ([]*domain.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "find")
	logOp.Infof("Request to find articles: %v", query)
	result, err := ctx.ArticleRepository().Find(ctx, query)
	if err != nil {
		errorMsg := fmt.Sprintf("error finding article with query: %v", query)
		return nil, errors.NewArticleError(errors.ArticleRetrievalErrorCode, errorMsg, err)
	}

	return result, err
}

func (u *articleUsecase) Create(ctx services.ArticleServiceContext, command *domain.ArticleCreateCommand) (*domain.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "create")
	logOp.Info("Request to create an article")

	if err := ctx.ArticleValidator().ValidateCreate(command); err != nil {
		errorMsg := fmt.Sprintf("error creating article with data: %v", command)
		return nil, errors.NewArticleError(errors.ArticleCreationErrorCode, errorMsg, err)
	}

	result, err := ctx.ArticleRepository().Create(ctx, command)
	if err != nil {
		errorMsg := fmt.Sprintf("error creating article with data: %v", command)
		return nil, errors.NewArticleError(errors.ArticleCreationErrorCode, errorMsg, err)
	}

	return result, nil
}

func (u *articleUsecase) Update(ctx services.ArticleServiceContext, command *domain.ArticleUpdateCommand) (*domain.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "update")
	logOp.Infof("Request to update an article: %d", command.ID)

	if err := ctx.ArticleValidator().ValidateUpdate(command); err != nil {
		errorMsg := fmt.Sprintf("error updating article with data: %v", command)
		return nil, errors.NewArticleError(errors.ArticleUpdateErrorCode, errorMsg, err)
	}

	result, err := ctx.ArticleRepository().Update(ctx, command)
	if err != nil {
		errorMsg := fmt.Sprintf("error updating article with data: %v", command)
		return nil, errors.NewArticleError(errors.ArticleUpdateErrorCode, errorMsg, err)
	}

	return result, nil
}

func (u *articleUsecase) Delete(ctx services.ArticleServiceContext, id uint) (*domain.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "delete")
	logOp.Infof("Request to delete an article: %d", id)

	result, err := ctx.ArticleRepository().FindByID(ctx, id)
	if err != nil {
		errorMsg := fmt.Sprintf("error deleting article with id: %d", id)
		return nil, errors.NewArticleError(errors.ArticleDeletionErrorCode, errorMsg, err)
	}

	if err = ctx.ArticleValidator().ValidateDelete(result); err != nil {
		errorMsg := fmt.Sprintf("error deleting article with id: %d", id)
		return nil, errors.NewArticleError(errors.ArticleDeletionErrorCode, errorMsg, err)
	}

	result, err = ctx.ArticleRepository().Delete(ctx, result.ID)
	if err != nil {
		errorMsg := fmt.Sprintf("error deleting article with id: %d", id)
		return nil, errors.NewArticleError(errors.ArticleDeletionErrorCode, errorMsg, err)
	}

	return result, nil
}

func (u *articleUsecase) Recover(ctx services.ArticleServiceContext, id uint) (*domain.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "recover")
	logOp.Infof("Request to recover an article: %d", id)

	result, err := ctx.ArticleRepository().FindDeletedByID(ctx, id)
	if err != nil {
		errorMsg := fmt.Sprintf("error recovering article with id: %d", id)
		return nil, errors.NewArticleError(errors.ArticleRecoveryErrorCode, errorMsg, err)
	}

	if err = ctx.ArticleValidator().ValidateRecover(result); err != nil {
		errorMsg := fmt.Sprintf("error recovering article with id: %d", id)
		return nil, errors.NewArticleError(errors.ArticleRecoveryErrorCode, errorMsg, err)
	}

	result, err = ctx.ArticleRepository().Recover(ctx, result.ID)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	return result, nil
}

func initLog(log *logger.Logger) *logrus.Entry {
	return log.WithFields(map[string]interface{}{"module": "service", "domain": "article"})
}
