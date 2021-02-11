package postgres

import "github.com/go-pg/pg/v10"

type Context interface {
	DB() *pg.DB
	Transaction() *pg.Tx
}
