package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/pkg/articles/ports/services"
)

type Article struct {
	ArticleService services.IArticle
}

func New(articleService services.IArticle) *Article {
	return &Article{ArticleService: articleService}
}

func (h *Article) Get(e echo.Context) error {
	return e.NoContent(http.StatusOK)
}
