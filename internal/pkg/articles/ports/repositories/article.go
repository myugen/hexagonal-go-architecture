package repositories

import (
	"context"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"
)

type IArticle interface {
	FindByID(ctx context.Context, id string) (*models.Article, error)
	Find(ctx context.Context, query *models.ArticleQuery) ([]*models.Article, error)
}
