package logger

import (
	"sync"

	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

var (
	logger *Logger
	once   sync.Once
)

func Initialize() {
	once.Do(func() {
		create()
	})
}

func Log() *Logger {
	return logger
}

func create() {
	logger = &Logger{logrus.New()}
	if viper.GetBool("verbose") {
		logger.SetLevel(log.DEBUG)
	} else {
		log.SetLevel(log.INFO)
	}
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
