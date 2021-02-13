package repositories

import (
	"github.com/myugen/hexagonal-go-architecture/context"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain/models"
)

type ArticleRepositoryContext interface {
	context.Context
}

type ArticleRepository interface {
	FindByID(ctx ArticleRepositoryContext, id uint) (*models.Article, error)
	FindDeletedByID(ctx ArticleRepositoryContext, id uint) (*models.Article, error)
	Find(ctx ArticleRepositoryContext, query *models.ArticleQuery) ([]*models.Article, error)
	Create(ctx ArticleRepositoryContext, command *models.ArticleCreateCommand) (*models.Article, error)
	Update(ctx ArticleRepositoryContext, command *models.ArticleUpdateCommand) (*models.Article, error)
	Delete(ctx ArticleRepositoryContext, id uint) (*models.Article, error)
	Recover(ctx ArticleRepositoryContext, id uint) (*models.Article, error)
}
