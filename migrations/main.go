package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"

	"github.com/mitchellh/go-homedir"
	"github.com/myugen/hexagonal-go-architecture/utils/constants"
	"github.com/spf13/viper"

	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	if viper.IsSet("mode") {
		return
	}

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Search config in home directory with name ".config" (without extension).
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", constants.AppLabel)) // path to look for the config file in
	viper.AddConfigPath(fmt.Sprintf("%s/.%s", home, constants.AppLabel))
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	viper.SetEnvPrefix(constants.AppLabel)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	dbConfig := viper.GetStringMapString("db")
	db := pg.Connect(&pg.Options{
		Addr:            fmt.Sprintf("%s:%s", dbConfig["host"], dbConfig["port"]),
		User:            dbConfig["user"],
		Database:        dbConfig["database"],
		Password:        dbConfig["password"],
		ApplicationName: constants.AppName,
		OnConnect: func(ctx context.Context, cn *pg.Conn) error {
			_, cnErr := cn.Exec(`SET TIME ZONE 'UTC';`)
			return cnErr
		},
		IdleTimeout: 30 * time.Second,
	})

	if _, err := db.Exec("SELECT 1"); err != nil {
		log.Fatalln(err)
	}

	if err := migrations.Run(db, constants.MigrationsDirectory, os.Args); err != nil {
		log.Fatalln(err)
	}
}
