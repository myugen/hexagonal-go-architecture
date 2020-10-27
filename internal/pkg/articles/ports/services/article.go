package services

import (
	"context"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/ports/repositories"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"

	"github.com/pkg/errors"
)

var (
	ErrExist = errors.New("Article already exist")
)

type IArticle interface {
	Get(ctx context.Context, id string) (*models.Article, error)
}

type article struct {
	articleRepository repositories.IArticle
}

func New(articleRepository repositories.IArticle) *article {
	return &article{articleRepository: articleRepository}
}

func (s *article) Get(ctx context.Context, id string) (*models.Article, error) {
	result, err := s.articleRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
