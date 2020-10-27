package repositories

import (
	"context"

	"github.com/myugen/hexagonal-go-architecture/pkg/articles/domain/models"
)

type IArticle interface {
	FindByID(ctx context.Context, id string) (*models.Article, error)
}
