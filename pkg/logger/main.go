package logger

import (
	"flag"
	"io/ioutil"
	"sync"

	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

var (
	logger *Logger
	once   sync.Once
)

func Log() *Logger {
	once.Do(func() {
		create()
	})
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

	if flag.Lookup("test.v") != nil {
		logger.Out = ioutil.Discard
	}
}
