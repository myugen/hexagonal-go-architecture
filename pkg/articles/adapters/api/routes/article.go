package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/pkg/articles/adapters/api/handlers"
)

func RegisterRoute(e *echo.Echo, articleHandler *handlers.Article) {
	article := e.Group("/articles")
	article.GET("/:id", articleHandler.Get)
}
