package routes

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/pkg/articles/adapters/api/handlers"
	"github.com/myugen/hexagonal-go-architecture/pkg/articles/adapters/db/dao"
	"github.com/myugen/hexagonal-go-architecture/pkg/articles/ports/services"
)

func RegisterRoute(e *echo.Echo) {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()

	articleRepository := dao.New(db)

	articleService := services.New(articleRepository)

	articleHandler := handlers.New(articleService)

	articleRoute := e.Group("/articles")
	articleRoute.GET("/:id", articleHandler.Get)
}
