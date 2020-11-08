package migrations

import (
	"github.com/go-pg/pg/v10"

	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func Migrate(db *pg.DB) error {
	return migrations.Run(db, "migrations", []string{"migrations/*.go", "migrate"})
}
