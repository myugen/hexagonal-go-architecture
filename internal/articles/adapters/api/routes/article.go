package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/api/context"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/api/handlers"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/usecase"
	"github.com/myugen/hexagonal-go-architecture/pkg/echo/middlewares"
	"github.com/myugen/hexagonal-go-architecture/pkg/logger"
	"github.com/myugen/hexagonal-go-architecture/pkg/postgres"
)

func RegisterRoute(parent *echo.Group) {

	articleRoute := parent.Group("/articles", ArticleAPIContextMiddleware(), middlewares.Transactional())

	articleHandler := handlers.New(usecase.NewArticleUsecase())
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
			cc := context.NewArticleAPIContext(c, postgres.DB(), logger.Log())
			return next(cc)
		}
	}
}
