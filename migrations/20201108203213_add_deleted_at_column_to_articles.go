package migrations

import (
	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			ALTER TABLE articles
				ADD COLUMN deleted_at TIMESTAMPTZ;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			ALTER TABLE articles
				DROP COLUMN deleted_at;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20201108203213_add_deleted_at_column_to_articles", up, down, opts)
}
