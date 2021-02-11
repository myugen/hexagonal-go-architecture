package dao

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain/models"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/repositories"
)

type ArticleDaoMock struct{}

func (a ArticleDaoMock) FindByID(ctx repositories.ArticleRepositoryContext, id uint) (*models.Article, error) {
	return nil, nil
}

func (a ArticleDaoMock) FindDeletedByID(ctx repositories.ArticleRepositoryContext, id uint) (*models.Article, error) {
	return nil, nil
}

func (a ArticleDaoMock) Find(ctx repositories.ArticleRepositoryContext, query *models.ArticleQuery) ([]*models.Article, error) {
	return nil, nil
}

func (a ArticleDaoMock) Create(ctx repositories.ArticleRepositoryContext, command *models.ArticleCreateCommand) (*models.Article, error) {
	return nil, nil
}

func (a ArticleDaoMock) Update(ctx repositories.ArticleRepositoryContext, command *models.ArticleUpdateCommand) (*models.Article, error) {
	return nil, nil
}

func (a ArticleDaoMock) Delete(ctx repositories.ArticleRepositoryContext, id uint) (*models.Article, error) {
	return nil, nil
}

func (a ArticleDaoMock) Recover(ctx repositories.ArticleRepositoryContext, id uint) (*models.Article, error) {
	return nil, nil
}

func NewArticleDaoMock() *ArticleDaoMock {
	return &ArticleDaoMock{}
}
