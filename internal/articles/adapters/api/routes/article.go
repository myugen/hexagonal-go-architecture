package routes

import (
	"github.com/labstack/echo/v4"
	commonMiddlewares "github.com/myugen/hexagonal-go-architecture/infrastructure/echo/middlewares"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/api/handlers"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/api/middlewares"
)

func RegisterRoute(parent *echo.Group, handler *handlers.ArticleHandler) {
	articleRoute := parent.Group("/articles",
		middlewares.ArticleAppContextMiddleware(),
		middlewares.ArticleErrorHandlerMiddleware(),
		commonMiddlewares.Transactional())

	articleRoute.GET("", handler.Find)
	articleRoute.GET("/:id", handler.Get)
	articleRoute.POST("", handler.Create)
	articleRoute.PUT("/:id", handler.Update)
	articleRoute.DELETE("/:id", handler.Delete)
	articleRoute.PATCH("/:id/recover", handler.Recover)
}
