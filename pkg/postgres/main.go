package postgres

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/myugen/hexagonal-go-architecture/utils/constants"

	"github.com/spf13/viper"

	"github.com/go-pg/pg/v10"
)

var (
	db    *pg.DB
	once  sync.Once
	mutex sync.Mutex
)

func Initialize() (err error) {
	once.Do(func() {
		err = create()
	})
	return err
}

func DB() *pg.DB {
	mutex.Lock()
	defer mutex.Unlock()

	return db
}

func Close() error {
	if db == nil {
		return nil
	}
	return db.Close()
}

func create() error {
	dbConfig := viper.GetStringMapString("db")
	db = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", dbConfig["host"], dbConfig["port"]),
		User:     dbConfig["user"],
		Database: dbConfig["postgres"],
		Password: dbConfig["password"],
		OnConnect: func(ctx context.Context, cn *pg.Conn) error {
			_, cnErr := cn.Exec(`SET application_name = ?; SET TIME ZONE 'UTC';`, constants.AppName)
			return cnErr
		},
		IdleTimeout: 30 * time.Second,
	})

	if _, err := db.Exec("SELECT 1"); err != nil {
		return err
	}

	return nil
}
