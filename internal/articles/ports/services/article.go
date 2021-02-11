package services

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain/models"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/usecase"
)

type IArticle interface {
	Get(ctx usecase.ArticleUsecaseContext, id uint) (*models.Article, error)
	Find(ctx usecase.ArticleUsecaseContext, query *models.ArticleQuery) ([]*models.Article, error)
	Create(ctx usecase.ArticleUsecaseContext, command *models.ArticleCreateCommand) (*models.Article, error)
	Update(ctx usecase.ArticleUsecaseContext, command *models.ArticleUpdateCommand) (*models.Article, error)
	Delete(ctx usecase.ArticleUsecaseContext, id uint) (*models.Article, error)
	Recover(ctx usecase.ArticleUsecaseContext, id uint) (*models.Article, error)
}
