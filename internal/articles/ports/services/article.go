package services

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain/models"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/repositories"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/validators"
)

type ArticleServiceContext interface {
	ArticleRepository() repositories.ArticleRepository
	ArticleValidator() validators.ArticleValidator
	repositories.ArticleRepositoryContext
}

type ArticleService interface {
	Get(ctx ArticleServiceContext, id uint) (*models.Article, error)
	Find(ctx ArticleServiceContext, query *models.ArticleQuery) ([]*models.Article, error)
	Create(ctx ArticleServiceContext, command *models.ArticleCreateCommand) (*models.Article, error)
	Update(ctx ArticleServiceContext, command *models.ArticleUpdateCommand) (*models.Article, error)
	Delete(ctx ArticleServiceContext, id uint) (*models.Article, error)
	Recover(ctx ArticleServiceContext, id uint) (*models.Article, error)
}
