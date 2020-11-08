package migrations

import (
	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE authors (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL
			);
			CREATE TABLE articles (
				id SERIAL PRIMARY KEY,
				title VARCHAR(255) NOT NULL,
				content TEXT NOT NULL,
				author INTEGER REFERENCES authors (id),
				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
			);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			DROP TABLE articles;
			DROP TABLE authors;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20201027235036_initial_schema", up, down, opts)
}
