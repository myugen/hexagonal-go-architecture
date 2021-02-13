package usecase

import (
	"github.com/go-pg/pg/v10"
	"github.com/myugen/hexagonal-go-architecture/infrastructure/logger"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/db/dao"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/validator"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/repositories"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/validators"
)

type ArticleServiceContextMock struct {
	repository repositories.ArticleRepository
	validator  validators.ArticleValidator
	logger     *logger.Logger
	db         *pg.DB
	tx         *pg.Tx
}

func (c *ArticleServiceContextMock) ArticleRepository() repositories.ArticleRepository {
	return c.repository
}

func (c *ArticleServiceContextMock) ArticleValidator() validators.ArticleValidator {
	return c.validator
}

func (c *ArticleServiceContextMock) Transaction() *pg.Tx {
	return c.tx
}

func (c *ArticleServiceContextMock) DB() *pg.DB {
	return c.db
}

func (c *ArticleServiceContextMock) Log() *logger.Logger {
	return c.logger
}

func NewArticleUsecaseContextMock() *ArticleServiceContextMock {
	return &ArticleServiceContextMock{
		repository: dao.NewArticleDaoMock(),
		validator:  validator.NewArticleValidator(),
		logger:     logger.Log(),
	}
}
