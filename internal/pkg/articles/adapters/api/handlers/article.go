package handlers

import (
	"net/http"
	"strconv"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/api/context"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/api/requests"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/api/responses"

	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/ports/services"
)

type IArticle interface {
	Get(e echo.Context) error
	Find(e echo.Context) error
	Create(e echo.Context) error
	Update(e echo.Context) error
	Delete(e echo.Context) error
	Recover(e echo.Context) error
}

type articleHandler struct {
	articleService services.IArticle
}

func New(articleService services.IArticle) *articleHandler {
	return &articleHandler{articleService: articleService}
}

func (h *articleHandler) Get(c echo.Context) error {
	ctx := c.(*context.ArticleAPIContext)
	param := c.Param("id")
	aux, err := strconv.ParseUint(param, 10, 64)
	id := uint(aux)
	article, err := h.articleService.Get(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs getting an article")
	}
	return c.JSON(http.StatusOK, responses.NewArticleResponse(article))
}

func (h *articleHandler) Find(c echo.Context) error {
	ctx := c.(*context.ArticleAPIContext)
	query := c.QueryParams()
	qpArticle := requests.NewArticleQueryParams(query)
	qArticle := qpArticle.ToArticleQuery()
	articles, err := h.articleService.Find(ctx, qArticle)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs getting articles")
	}

	response := make([]*responses.ArticleResponse, 0)
	for _, article := range articles {
		response = append(response, responses.NewArticleResponse(article))
	}

	return c.JSON(http.StatusOK, response)
}

func (h *articleHandler) Create(c echo.Context) error {
	ctx := c.(*context.ArticleAPIContext)
	body := new(requests.ArticleCreateRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs creating an article: %s")
	}

	article, err := h.articleService.Create(ctx, body.ToCommand())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs creating an article")
	}

	return c.JSON(http.StatusOK, responses.NewArticleResponse(article))
}

func (h *articleHandler) Update(c echo.Context) error {
	ctx := c.(*context.ArticleAPIContext)
	id := c.Param("id")
	body := new(requests.ArticleUpdateRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs updating an article")
	}
	aux, err := strconv.ParseUint(id, 10, 64)
	body.ID = uint(aux)

	article, err := h.articleService.Update(ctx, body.ToCommand())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs updating an article", err)
	}

	return c.JSON(http.StatusOK, responses.NewArticleResponse(article))
}

func (h *articleHandler) Delete(c echo.Context) error {
	ctx := c.(*context.ArticleAPIContext)
	param := c.Param("id")
	aux, err := strconv.ParseUint(param, 10, 64)
	id := uint(aux)
	article, err := h.articleService.Delete(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs deleting an article")
	}
	return c.JSON(http.StatusOK, responses.NewArticleResponse(article))
}

func (h *articleHandler) Recover(c echo.Context) error {
	ctx := c.(*context.ArticleAPIContext)
	param := c.Param("id")
	aux, err := strconv.ParseUint(param, 10, 64)
	id := uint(aux)
	article, err := h.articleService.Recover(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs recovering an article")
	}
	return c.JSON(http.StatusOK, responses.NewArticleResponse(article))
}
