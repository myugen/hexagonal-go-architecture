package context

import "github.com/go-pg/pg/v10"

type Context interface {
	Transaction() *pg.Tx
	DB() *pg.DB
}
