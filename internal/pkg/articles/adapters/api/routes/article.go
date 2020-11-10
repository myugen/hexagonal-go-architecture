package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/api/context"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/api/handlers"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/ports/services"
	"github.com/myugen/hexagonal-go-architecture/pkg/echo/middlewares"
	"github.com/myugen/hexagonal-go-architecture/pkg/postgres"
)

func RegisterRoute(parent *echo.Group) {

	articleRoute := parent.Group("/articles", ArticleAPIContextMiddleware(), middlewares.Transactional())

	articleHandler := handlers.New(services.New())
	articleRoute.GET("", articleHandler.Find)
	articleRoute.GET("/:id", articleHandler.Get)
	articleRoute.POST("", articleHandler.Create)
	articleRoute.PUT("/:id", articleHandler.Update)
	articleRoute.DELETE("/:id", articleHandler.Delete)
	articleRoute.PATCH("/:id/recover", articleHandler.Recover)
}

func ArticleAPIContextMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := context.NewArticleAPIContext(c, postgres.DB())
			return next(cc)
		}
	}
}
