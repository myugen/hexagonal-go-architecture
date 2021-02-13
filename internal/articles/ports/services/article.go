package services

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/repositories"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/validators"
)

type ArticleServiceContext interface {
	ArticleRepository() repositories.ArticleRepository
	ArticleValidator() validators.ArticleValidator
	repositories.ArticleRepositoryContext
}

type ArticleService interface {
	Get(ctx ArticleServiceContext, id uint) (*domain.Article, error)
	Find(ctx ArticleServiceContext, query *domain.ArticleQuery) ([]*domain.Article, error)
	Create(ctx ArticleServiceContext, command *domain.ArticleCreateCommand) (*domain.Article, error)
	Update(ctx ArticleServiceContext, command *domain.ArticleUpdateCommand) (*domain.Article, error)
	Delete(ctx ArticleServiceContext, id uint) (*domain.Article, error)
	Recover(ctx ArticleServiceContext, id uint) (*domain.Article, error)
}
