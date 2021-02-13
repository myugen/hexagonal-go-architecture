package app

import (
	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/api/handlers"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/api/routes"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/usecase"
)

func SetupArticleApp(parent *echo.Group) {
	handler := handlers.NewArticleHandler(usecase.NewArticleUsecase())
	routes.RegisterRoute(parent, handler)
}
