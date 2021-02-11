package middlewares

import (
	"errors"

	e "github.com/myugen/hexagonal-go-architecture/infrastructure/echo"

	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/utils/constants"
)

const ErrTransactionerNotImplemented = "middleware: transactioner needs to be implemented"

func Transactional() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.(e.Context)
			if !ok {
				return errors.New(ErrTransactionerNotImplemented)
			}

			tx, _ := ctx.BeginTransaction()
			c.Set(constants.TxKey, tx)

			if err := next(c); err != nil {
				ctx.RollbackTransaction()
				return err
			}
			ctx.CommitTransaction()

			return nil
		}
	}
}
