package middlewares

import (
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/utils/constants"
)

const ErrTransactionerNotImplemented = "middleware: transactioner needs to be implemented"

type Transactioner interface {
	BeginTransaction() (*pg.Tx, error)
}

func Transactional() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.(Transactioner)
			if !ok {
				return errors.New(ErrTransactionerNotImplemented)
			}

			tx, _ := ctx.BeginTransaction()
			c.Set(constants.TxKey, tx)

			if err := next(c); err != nil {
				tx.Rollback()
				return err
			}
			tx.Commit()

			return nil
		}
	}
}
