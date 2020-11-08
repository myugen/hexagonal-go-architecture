package services

import (
	"context"

	"github.com/go-playground/validator/v10"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/ports/repositories"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"

	"github.com/pkg/errors"
)

var (
	errAlreadyDeleted = errors.New("article was already deleted")
)

type IArticle interface {
	Get(ctx context.Context, id uint) (*models.Article, error)
	Find(ctx context.Context, query *models.ArticleQuery) ([]*models.Article, error)
	Create(ctx context.Context, command *models.ArticleCreateCommand) (*models.Article, error)
	Update(ctx context.Context, command *models.ArticleUpdateCommand) (*models.Article, error)
	Delete(ctx context.Context, id uint) (*models.Article, error)
}

type article struct {
	articleRepository repositories.IArticle
	validate          *validator.Validate
}

func New(articleRepository repositories.IArticle) *article {
	return &article{articleRepository: articleRepository, validate: validator.New()}
}

func (s *article) Get(ctx context.Context, id uint) (*models.Article, error) {
	result, err := s.articleRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *article) Find(ctx context.Context, query *models.ArticleQuery) ([]*models.Article, error) {
	result, err := s.articleRepository.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (s *article) Create(ctx context.Context, command *models.ArticleCreateCommand) (*models.Article, error) {
	if err := s.validate.Struct(command); err != nil {
		return nil, err
	}

	result, err := s.articleRepository.Create(ctx, command)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *article) Update(ctx context.Context, command *models.ArticleUpdateCommand) (*models.Article, error) {
	if err := s.validate.Struct(command); err != nil {
		return nil, err
	}

	result, err := s.articleRepository.Update(ctx, command)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *article) Delete(ctx context.Context, id uint) (*models.Article, error) {
	result, err := s.articleRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if result.IsDeleted {
		return nil, errAlreadyDeleted
	}

	result, err = s.articleRepository.Delete(ctx, result.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
