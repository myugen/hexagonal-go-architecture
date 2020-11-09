package services

import (
	"context"

	"github.com/myugen/hexagonal-go-architecture/pkg/logger"
	"github.com/sirupsen/logrus"

	"github.com/go-playground/validator/v10"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/ports/repositories"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"

	"github.com/pkg/errors"
)

var (
	errAlreadyDeleted = errors.New("article was already deleted")
	errNotDeleted     = errors.New("article is not deleted")
)

var log *logrus.Entry

type IArticle interface {
	Get(ctx context.Context, id uint) (*models.Article, error)
	Find(ctx context.Context, query *models.ArticleQuery) ([]*models.Article, error)
	Create(ctx context.Context, command *models.ArticleCreateCommand) (*models.Article, error)
	Update(ctx context.Context, command *models.ArticleUpdateCommand) (*models.Article, error)
	Delete(ctx context.Context, id uint) (*models.Article, error)
	Recover(ctx context.Context, id uint) (*models.Article, error)
}

type article struct {
	articleRepository repositories.IArticle
	validate          *validator.Validate
}

func New(articleRepository repositories.IArticle) *article {
	log = logger.Log().WithFields(map[string]interface{}{"module": "service", "domain": "article"})
	return &article{articleRepository: articleRepository, validate: validator.New()}
}

func (s *article) Get(ctx context.Context, id uint) (*models.Article, error) {
	logOp := log.WithField("operation", "get")
	logOp.Infof("Request to get an article: %d", id)
	result, err := s.articleRepository.FindByID(ctx, id)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}
	return result, nil
}

func (s *article) Find(ctx context.Context, query *models.ArticleQuery) ([]*models.Article, error) {
	logOp := log.WithField("operation", "find")
	logOp.Infof("Request to find articles: %v", query)
	result, err := s.articleRepository.Find(ctx, query)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	return result, err
}

func (s *article) Create(ctx context.Context, command *models.ArticleCreateCommand) (*models.Article, error) {
	logOp := log.WithField("operation", "create")
	logOp.Info("Request to create an article")

	if err := s.validate.Struct(command); err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	result, err := s.articleRepository.Create(ctx, command)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	return result, nil
}

func (s *article) Update(ctx context.Context, command *models.ArticleUpdateCommand) (*models.Article, error) {
	logOp := log.WithField("operation", "update")
	logOp.Infof("Request to update an article: %d", command.ID)

	if err := s.validate.Struct(command); err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	result, err := s.articleRepository.Update(ctx, command)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	return result, nil
}

func (s *article) Delete(ctx context.Context, id uint) (*models.Article, error) {
	logOp := log.WithField("operation", "delete")
	logOp.Infof("Request to delete an article: %d", id)

	result, err := s.articleRepository.FindByID(ctx, id)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}
	if result.IsDeleted {
		logOp.WithField("error", errAlreadyDeleted).Errorf("article service error")
		return nil, errAlreadyDeleted
	}

	result, err = s.articleRepository.Delete(ctx, result.ID)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	return result, nil
}

func (s *article) Recover(ctx context.Context, id uint) (*models.Article, error) {
	logOp := log.WithField("operation", "recover")
	logOp.Infof("Request to recover an article: %d", id)

	result, err := s.articleRepository.FindDeletedByID(ctx, id)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}
	if !result.IsDeleted {
		logOp.WithField("error", errNotDeleted).Errorf("article service error")
		return nil, errNotDeleted
	}

	result, err = s.articleRepository.Recover(ctx, result.ID)
	if err != nil {
		logOp.WithField("error", err).Errorf("article service error")
		return nil, err
	}

	return result, nil
}
