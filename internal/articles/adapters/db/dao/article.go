package dao

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/repositories"

	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/db/entities"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/db/types"

	"github.com/go-pg/pg/v10"
)

type article struct{}

func NewArticleDAO() *article {
	return &article{}
}

func (d *article) FindByID(ctx repositories.ArticleRepositoryContext, id uint) (*domain.Article, error) {
	eArticles := new(entities.ArticleEntity)
	eArticles.ID = id
	if err := ctx.Transaction().Model(eArticles).
		WherePK().
		Relation("Author").
		Select(); err != nil {
		return nil, err
	}

	return eArticles.ToModel(), nil
}

func (d *article) FindDeletedByID(ctx repositories.ArticleRepositoryContext, id uint) (*domain.Article, error) {
	eArticles := new(entities.ArticleEntity)
	eArticles.ID = id
	if err := ctx.Transaction().Model(eArticles).
		WherePK().
		AllWithDeleted().
		Relation("Author").
		Select(); err != nil {
		return nil, err
	}

	return eArticles.ToModel(), nil
}

func (d *article) Find(ctx repositories.ArticleRepositoryContext, query *domain.ArticleQuery) ([]*domain.Article, error) {
	eArticles := new([]*entities.ArticleEntity)
	fArticle := types.NewArticleFilter(query)
	if err := ctx.Transaction().Model(eArticles).
		Relation("Author").
		Apply(fArticle.Where).
		Select(); err != nil {
		return nil, err
	}

	var articles []*domain.Article
	for _, eArticle := range *eArticles {
		articles = append(articles, eArticle.ToModel())
	}

	return articles, nil
}

func (d *article) Create(ctx repositories.ArticleRepositoryContext, command *domain.ArticleCreateCommand) (*domain.Article, error) {
	eArticle := entities.NewArticleEntity(command)
	if _, err := ctx.Transaction().Model(eArticle).
		Insert(); err != nil {
		return nil, err
	}

	return eArticle.ToModel(), nil
}

func (d *article) Update(ctx repositories.ArticleRepositoryContext, command *domain.ArticleUpdateCommand) (*domain.Article, error) {
	eArticle := new(entities.ArticleEntity)
	eArticle.ID = command.ID
	if err := ctx.Transaction().Model(eArticle).
		Relation("Author").
		WherePK().
		Select(); err != nil {
		return nil, err
	}

	eArticle.UpdateFrom(command)
	if _, err := ctx.Transaction().Model(eArticle).
		WherePK().
		UpdateNotZero(); err != nil {
		return nil, err
	}

	return eArticle.ToModel(), nil
}

func (d *article) Delete(ctx repositories.ArticleRepositoryContext, id uint) (*domain.Article, error) {
	eArticle := new(entities.ArticleEntity)
	eArticle.ID = id
	if _, err := ctx.Transaction().Model(eArticle).
		Relation("Author").
		WherePK().
		Delete(); err != nil {
		return nil, err
	}

	return eArticle.ToModel(), nil
}

func (d *article) Recover(ctx repositories.ArticleRepositoryContext, id uint) (*domain.Article, error) {
	eArticle := new(entities.ArticleEntity)
	eArticle.ID = id
	if err := ctx.Transaction().Model(eArticle).
		Relation("Author").
		WherePK().
		AllWithDeleted().
		Select(); err != nil {
		return nil, err
	}

	eArticle.DeletedAt = pg.NullTime{}
	if _, err := ctx.Transaction().Model(eArticle).
		WherePK().
		AllWithDeleted().
		Update(); err != nil {
		return nil, err
	}

	return eArticle.ToModel(), nil
}
