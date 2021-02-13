package repositories

import (
	"github.com/myugen/hexagonal-go-architecture/context"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain"
)

type ArticleRepositoryContext interface {
	context.Context
}

type ArticleRepository interface {
	FindByID(ctx ArticleRepositoryContext, id uint) (*domain.Article, error)
	FindDeletedByID(ctx ArticleRepositoryContext, id uint) (*domain.Article, error)
	Find(ctx ArticleRepositoryContext, query *domain.ArticleQuery) ([]*domain.Article, error)
	Create(ctx ArticleRepositoryContext, command *domain.ArticleCreateCommand) (*domain.Article, error)
	Update(ctx ArticleRepositoryContext, command *domain.ArticleUpdateCommand) (*domain.Article, error)
	Delete(ctx ArticleRepositoryContext, id uint) (*domain.Article, error)
	Recover(ctx ArticleRepositoryContext, id uint) (*domain.Article, error)
}
