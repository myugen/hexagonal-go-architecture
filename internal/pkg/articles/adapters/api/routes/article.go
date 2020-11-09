package routes

import (
	"github.com/myugen/hexagonal-go-architecture/pkg/postgres"

	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/api/handlers"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/db/dao"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/ports/services"
)

func RegisterRoute(parent *echo.Group) {

	articleRepository := dao.New(postgres.DB())

	articleService := services.New(articleRepository)

	articleHandler := handlers.New(articleService)

	articleRoute := parent.Group("/articles")
	articleRoute.GET("", articleHandler.Find)
	articleRoute.GET("/:id", articleHandler.Get)
	articleRoute.POST("", articleHandler.Create)
	articleRoute.PUT("/:id", articleHandler.Update)
	articleRoute.DELETE("/:id", articleHandler.Delete)
	articleRoute.PATCH("/:id/recover", articleHandler.Recover)
}
