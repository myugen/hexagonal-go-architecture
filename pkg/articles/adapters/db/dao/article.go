package dao

import (
	"context"

	"github.com/myugen/hexagonal-go-architecture/pkg/articles/adapters/db/entities"

	"github.com/go-pg/pg/v10"
	"github.com/myugen/hexagonal-go-architecture/pkg/articles/domain/models"
)

type article struct {
	db *pg.DB
}

func New(db *pg.DB) *article {
	return &article{db: db}
}

func (d *article) FindByID(ctx context.Context, id string) (*models.Article, error) {
	entity := new(entities.Article)
	if err := d.db.Model(entity).Where("article.id = ?", id).Relation("Author").Select(); err != nil {
		return nil, err
	}

	return entity.ToModel(), nil
}
