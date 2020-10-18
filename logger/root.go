package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var (
	once   sync.Once
	logger *Logger
)

func New() *Logger {
	once.Do(func() {
		logger = create()
	})
	return logger
}

func GetInstance() *Logger {
	return logger
}

func create() *Logger {
	return &Logger{Logger: logrus.New()}
}
