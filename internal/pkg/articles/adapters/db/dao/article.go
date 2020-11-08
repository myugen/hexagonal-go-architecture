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

func (d *article) FindByID(ctx context.Context, id uint) (*models.Article, error) {
	eArticles := new(entities.ArticleEntity)
	eArticles.ID = id
	if err := d.db.Model(eArticles).
		WherePK().
		Relation("Author").
		Select(); err != nil {
		return nil, err
	}

	return eArticles.ToModel(), nil
}

func (d *article) Find(ctx context.Context, query *models.ArticleQuery) ([]*models.Article, error) {
	eArticles := new([]*entities.ArticleEntity)
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

func (d *article) Create(ctx context.Context, command *models.ArticleCreateCommand) (*models.Article, error) {
	eArticle := entities.NewArticleEntity(command)
	if _, err := d.db.Model(eArticle).
		Insert(); err != nil {
		return nil, err
	}

	return eArticle.ToModel(), nil
}

func (d *article) Update(ctx context.Context, command *models.ArticleUpdateCommand) (*models.Article, error) {
	eArticle := new(entities.ArticleEntity)
	eArticle.ID = command.ID
	if err := d.db.Model(eArticle).
		Relation("Author").
		WherePK().
		Select(); err != nil {
		return nil, err
	}

	eArticle.UpdateFrom(command)
	if _, err := d.db.Model(eArticle).
		WherePK().
		UpdateNotZero(); err != nil {
		return nil, err
	}

	return eArticle.ToModel(), nil
}

func (d *article) Delete(ctx context.Context, id uint) (*models.Article, error) {
	eArticle := new(entities.ArticleEntity)
	eArticle.ID = id
	if _, err := d.db.Model(eArticle).
		Relation("Author").
		WherePK().
		Delete(); err != nil {
		return nil, err
	}

	return eArticle.ToModel(), nil
}
