package config

import (
	"sync"
)

var (
	config *AppConfig
	once   sync.Once
)

func Initialize(initialConfig AppConfig) {
	once.Do(func() {
		config = &initialConfig
	})
}

func Config() *AppConfig {
	return config
}

type AppConfig struct {
	Mode    string
	Version string
	Server  ServerConfiguration
	DB      DatabaseConfiguration
	Verbose bool
}

type ServerConfiguration struct {
	Port int
}

type DatabaseConfiguration struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}
