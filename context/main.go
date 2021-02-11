package context

import (
	"github.com/myugen/hexagonal-go-architecture/infrastructure/logger"
	"github.com/myugen/hexagonal-go-architecture/infrastructure/postgres"
)

type Context interface {
	postgres.Context
	logger.Context
}
