package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/infrastructure/logger"
	"github.com/myugen/hexagonal-go-architecture/infrastructure/postgres"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/app/context"
)

func ArticleAppContextMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := context.NewArticleAppContext(c, postgres.DB(), logger.Log())
			return next(cc)
		}
	}
}
