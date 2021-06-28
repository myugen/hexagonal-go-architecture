package errors

import (
	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/infrastructure/errors"
	"github.com/myugen/hexagonal-go-architecture/infrastructure/logger"
)

func HTTPErrorHandler(err error, _ echo.Context) {
	log := logger.Log()
	if contextor, ok := err.(errors.Contextor); ok {
		log.WithFields(contextor.Context()).Error(err.Error())
	} else {
		log.Error(err.Error())
	}
}
