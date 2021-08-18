package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/config"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/logger"

	"github.com/myugen/hexagonal-go-architecture/migrations"

	"github.com/pkg/errors"

	"github.com/myugen/hexagonal-go-architecture/utils/constants"

	"github.com/spf13/viper"

	"github.com/go-pg/pg/v10"
)

var (
	db   *pg.DB
	once sync.Once
)

func Initialize() (err error) {
	once.Do(func() {
		err = create()
	})
	return err
}

func DB() *pg.DB {
	return db
}

func Close() error {
	if db == nil {
		return nil
	}
	return db.Close()
}

func create() error {
	appConfig := config.Config()
	db = pg.Connect(&pg.Options{
		Addr:            fmt.Sprintf("%s:%d", appConfig.DB.Host, appConfig.DB.Port),
		User:            appConfig.DB.User,
		Database:        appConfig.DB.Database,
		Password:        appConfig.DB.Password,
		ApplicationName: constants.AppName,
		MaxRetries:      5,
		OnConnect: func(ctx context.Context, cn *pg.Conn) error {
			_, cnErr := cn.Exec(`SET TIME ZONE 'UTC';`)
			return cnErr
		},
	})

	verbose := viper.GetBool("verbose")
	db.AddQueryHook(NewLoggerHook(logger.Log(), verbose))

	if _, err := db.Exec("SELECT 1"); err != nil {
		return errors.Wrap(err, "database connection failure")
	}

	if err := migrations.Migrate(db); err != nil {
		return errors.Wrap(err, "database migration failure")
	}

	return nil
}
