package context

import (
	"github.com/go-pg/pg/v10"
	"github.com/myugen/hexagonal-go-architecture/pkg/logger"
)

type Context interface {
	Transaction() *pg.Tx
	DB() *pg.DB
	Log() *logger.Logger
}
