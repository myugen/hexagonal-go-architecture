package context

import (
	"github.com/myugen/hexagonal-go-architecture/pkg/logger"
	"github.com/myugen/hexagonal-go-architecture/pkg/postgres"
)

type Context interface {
	postgres.Context
	logger.Context
}
