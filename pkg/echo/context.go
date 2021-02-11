package echo

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
)

type Context interface {
	echo.Context
	BeginTransaction() (*pg.Tx, error)
	CommitTransaction() error
	RollbackTransaction() error
}
