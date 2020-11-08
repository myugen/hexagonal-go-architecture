package dao

import (
	"context"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/db/entities"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/db/types"

	"github.com/go-pg/pg/v10"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"
)

type article struct {
	db *pg.DB
}

func New(db *pg.DB) *article {
	return &article{db: db}
}

func (d *article) FindByID(ctx context.Context, id string) (*models.Article, error) {
	eArticles := new(entities.Article)
	if err := d.db.Model(eArticles).
		Where("article.id = ?", id).
		Relation("Author").
		Select(); err != nil {
		return nil, err
	}

	return eArticles.ToModel(), nil
}

func (d *article) Find(ctx context.Context, query *models.ArticleQuery) ([]*models.Article, error) {
	eArticles := new([]*entities.Article)
	fArticle := types.NewArticleFilter(query)
	if err := d.db.Model(eArticles).
		Relation("Author").
		Apply(fArticle.Where).
		Select(); err != nil {
		return nil, err
	}

	var articles []*models.Article
	for _, eArticle := range *eArticles {
		articles = append(articles, eArticle.ToModel())
	}

	return articles, nil
}
