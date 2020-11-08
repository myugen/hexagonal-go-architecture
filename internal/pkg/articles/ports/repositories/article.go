package repositories

import (
	"context"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"
)

type IArticle interface {
	FindByID(ctx context.Context, id string) (*models.Article, error)
	Find(ctx context.Context, query *models.ArticleQuery) ([]*models.Article, error)
	Create(ctx context.Context, command *models.ArticleCreateCommand) (*models.Article, error)
	Update(ctx context.Context, command *models.ArticleUpdateCommand) (*models.Article, error)
}
