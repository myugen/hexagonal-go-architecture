package services

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain/models"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/usecase"
)

type IArticle interface {
	Get(ctx usecase.ArticleContext, id uint) (*models.Article, error)
	Find(ctx usecase.ArticleContext, query *models.ArticleQuery) ([]*models.Article, error)
	Create(ctx usecase.ArticleContext, command *models.ArticleCreateCommand) (*models.Article, error)
	Update(ctx usecase.ArticleContext, command *models.ArticleUpdateCommand) (*models.Article, error)
	Delete(ctx usecase.ArticleContext, id uint) (*models.Article, error)
	Recover(ctx usecase.ArticleContext, id uint) (*models.Article, error)
}
