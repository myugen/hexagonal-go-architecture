package usecase

import (
	"github.com/go-pg/pg/v10"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/db/dao"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/validator"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/repositories"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/validators"
	"github.com/myugen/hexagonal-go-architecture/pkg/logger"
)

type ArticleUsecaseContextMock struct {
	repository repositories.IArticle
	validator  validators.IArticle
	logger     *logger.Logger
	db         *pg.DB
	tx         *pg.Tx
}

func (c *ArticleUsecaseContextMock) ArticleRepository() repositories.IArticle {
	return c.repository
}

func (c *ArticleUsecaseContextMock) ArticleValidator() validators.IArticle {
	return c.validator
}

func (c *ArticleUsecaseContextMock) Transaction() *pg.Tx {
	return c.tx
}

func (c *ArticleUsecaseContextMock) DB() *pg.DB {
	return c.db
}

func (c *ArticleUsecaseContextMock) Log() *logger.Logger {
	return c.logger
}

func NewArticleUsecaseContextMock() *ArticleUsecaseContextMock {
	return &ArticleUsecaseContextMock{
		repository: dao.NewArticleDaoMock(),
		validator:  validator.NewArticleValidator(),
		logger:     logger.Log(),
	}
}
