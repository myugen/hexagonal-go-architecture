package context

import (
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type Context interface {
	Transaction() *pg.Tx
	DB() *pg.DB
	Log() *logrus.Logger
}
