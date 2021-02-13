package context

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/infrastructure/logger"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/db/dao"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/validator"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/repositories"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/validators"
	"github.com/pkg/errors"
)

type ArticleAppContext struct {
	echo.Context
	db                *pg.DB
	tx                *pg.Tx
	log               *logger.Logger
	articleRepository repositories.ArticleRepository
	articleValidator  validators.ArticleValidator
}

func NewArticleAppContext(c echo.Context, db *pg.DB, log *logger.Logger) *ArticleAppContext {
	return &ArticleAppContext{
		Context:           c,
		db:                db,
		articleRepository: dao.NewArticleDAO(),
		log:               log,
		articleValidator:  validator.NewArticleValidator(),
	}
}

func (a *ArticleAppContext) ArticleRepository() repositories.ArticleRepository {
	return a.articleRepository
}

func (a *ArticleAppContext) ArticleValidator() validators.ArticleValidator {
	return a.articleValidator
}

func (a *ArticleAppContext) Transaction() *pg.Tx {
	return a.tx
}

func (a *ArticleAppContext) DB() *pg.DB {
	return a.db
}

func (a *ArticleAppContext) Log() *logger.Logger {
	return a.log
}

func (a *ArticleAppContext) CommitTransaction() error {
	return a.tx.Commit()
}
func (a *ArticleAppContext) RollbackTransaction() error {
	return a.tx.Rollback()
}

func (a *ArticleAppContext) BeginTransaction() (*pg.Tx, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	a.tx = tx
	return a.tx, nil
}
