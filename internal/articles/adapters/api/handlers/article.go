package handlers

import (
	"net/http"
	"strconv"

	"github.com/myugen/hexagonal-go-architecture/internal/articles/app/context"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/decoders"

	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/api/requests"

	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/api/responses"

	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/services"
)

type ArticleHandler struct {
	articleService services.ArticleService
}

func NewArticleHandler(articleService services.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (h *ArticleHandler) Get(c echo.Context) error {
	ctx := c.(*context.ArticleAppContext)
	param := ctx.Param("id")
	aux, err := strconv.ParseUint(param, 10, 64)
	id := uint(aux)
	article, err := h.articleService.Get(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs getting an article")
	}
	return ctx.JSON(http.StatusOK, responses.NewArticleResponse(article))
}

func (h *ArticleHandler) Find(c echo.Context) error {
	ctx := c.(*context.ArticleAppContext)
	decoder := decoders.ParamsDecoder()
	qpArticle := new(requests.ArticleQueryParams)
	if err := decoder.Decode(qpArticle, ctx.QueryParams()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs getting articles")
	}
	qArticle := qpArticle.ToArticleQuery()
	articles, err := h.articleService.Find(ctx, qArticle)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs getting articles")
	}

	response := make([]*responses.ArticleResponse, 0)
	for _, article := range articles {
		response = append(response, responses.NewArticleResponse(article))
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *ArticleHandler) Create(c echo.Context) error {
	ctx := c.(*context.ArticleAppContext)
	body := new(requests.ArticleCreateRequest)
	if err := ctx.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs creating an article: %s")
	}

	article, err := h.articleService.Create(ctx, body.ToCommand())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs creating an article")
	}

	return ctx.JSON(http.StatusOK, responses.NewArticleResponse(article))
}

func (h *ArticleHandler) Update(c echo.Context) error {
	ctx := c.(*context.ArticleAppContext)
	id := ctx.Param("id")
	body := new(requests.ArticleUpdateRequest)
	if err := ctx.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs updating an article")
	}
	aux, err := strconv.ParseUint(id, 10, 64)
	body.ID = uint(aux)

	article, err := h.articleService.Update(ctx, body.ToCommand())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs updating an article", err)
	}

	return ctx.JSON(http.StatusOK, responses.NewArticleResponse(article))
}

func (h *ArticleHandler) Delete(c echo.Context) error {
	ctx := c.(*context.ArticleAppContext)
	param := ctx.Param("id")
	aux, err := strconv.ParseUint(param, 10, 64)
	id := uint(aux)
	article, err := h.articleService.Delete(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs deleting an article")
	}
	return ctx.JSON(http.StatusOK, responses.NewArticleResponse(article))
}

func (h *ArticleHandler) Recover(c echo.Context) error {
	ctx := c.(*context.ArticleAppContext)
	param := ctx.Param("id")
	aux, err := strconv.ParseUint(param, 10, 64)
	id := uint(aux)
	article, err := h.articleService.Recover(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs recovering an article")
	}
	return ctx.JSON(http.StatusOK, responses.NewArticleResponse(article))
}
