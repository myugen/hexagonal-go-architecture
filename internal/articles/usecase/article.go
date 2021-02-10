package usecase

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/repositories"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/validators"
	"github.com/myugen/hexagonal-go-architecture/pkg/logger"

	"github.com/sirupsen/logrus"

	"github.com/go-playground/validator/v10"

	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain/models"
)

type ArticleContext interface {
	ArticleRepository() repositories.IArticle
	ArticleValidator() validators.IArticle
	repositories.ArticleRepositoryContext
}

type articleUsecase struct {
	validate *validator.Validate
}

func NewArticleUsecase() *articleUsecase {
	return &articleUsecase{validate: validator.New()}
}

func (u *articleUsecase) Get(ctx ArticleContext, id uint) (*models.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "get")
	logOp.Infof("Request to get an article: %d", id)
	result, err := ctx.ArticleRepository().FindByID(ctx, id)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}
	return result, nil
}

func (u *articleUsecase) Find(ctx ArticleContext, query *models.ArticleQuery) ([]*models.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "find")
	logOp.Infof("Request to find articles: %v", query)
	result, err := ctx.ArticleRepository().Find(ctx, query)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	return result, err
}

func (u *articleUsecase) Create(ctx ArticleContext, command *models.ArticleCreateCommand) (*models.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "create")
	logOp.Info("Request to create an article")

	if err := ctx.ArticleValidator().ValidateCreate(command); err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	result, err := ctx.ArticleRepository().Create(ctx, command)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	return result, nil
}

func (u *articleUsecase) Update(ctx ArticleContext, command *models.ArticleUpdateCommand) (*models.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "update")
	logOp.Infof("Request to update an article: %d", command.ID)

	if err := ctx.ArticleValidator().ValidateUpdate(command); err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	result, err := ctx.ArticleRepository().Update(ctx, command)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	return result, nil
}

func (u *articleUsecase) Delete(ctx ArticleContext, id uint) (*models.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "delete")
	logOp.Infof("Request to delete an article: %d", id)

	result, err := ctx.ArticleRepository().FindByID(ctx, id)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}
	if err = ctx.ArticleValidator().ValidateDelete(result); err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	result, err = ctx.ArticleRepository().Delete(ctx, result.ID)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	return result, nil
}

func (u *articleUsecase) Recover(ctx ArticleContext, id uint) (*models.Article, error) {
	logOp := initLog(ctx.Log()).WithField("operation", "recover")
	logOp.Infof("Request to recover an article: %d", id)

	result, err := ctx.ArticleRepository().FindDeletedByID(ctx, id)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	if err = ctx.ArticleValidator().ValidateRecover(result); err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
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
