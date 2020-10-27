package logger

import (
	"sync"

	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

var (
	log  *logrus.Logger
	once sync.Once
)

func Initialize() {
	once.Do(func() {
		create()
	})
}

func Log() *logrus.Logger {
	return log
}

func create() {
	log = logrus.New()
	if viper.GetBool("verbose") {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
