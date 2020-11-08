package requests

import (
	"net/url"
	"strconv"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"
)

type ArticleCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r *ArticleCreateRequest) ToCommand() *models.ArticleCreateCommand {
	return &models.ArticleCreateCommand{
		Title:   r.Title,
		Content: r.Content,
	}
}

type ArticleUpdateRequest struct {
	ID uint `json:"id"`
	*ArticleCreateRequest
}

func (r *ArticleUpdateRequest) ToCommand() *models.ArticleUpdateCommand {
	return &models.ArticleUpdateCommand{
		ID: r.ID,
		ArticleCreateCommand: &models.ArticleCreateCommand{
			Title:   r.Title,
			Content: r.Content,
		},
	}
}

type ArticleQueryParams struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	AuthorID uint   `json:"author_id"`
	Title    string `json:"title"`
}

func NewArticleQueryParams(values url.Values) *ArticleQueryParams {
	qpArticle := new(ArticleQueryParams)
	offset := values.Get("offset")
	if offset != "" {
		qpArticle.Offset, _ = strconv.Atoi(offset)
	}

	limit := values.Get("limit")
	if limit != "" {
		qpArticle.Limit, _ = strconv.Atoi(limit)
	}

	authorID := values.Get("author_id")
	if authorID != "" {
		aux, _ := strconv.ParseUint(authorID, 10, 32)
		qpArticle.AuthorID = uint(aux)
	}

	qpArticle.Title = values.Get("title")

	return qpArticle
}

func (a *ArticleQueryParams) ToArticleQuery() *models.ArticleQuery {
	return &models.ArticleQuery{
		Offset:   a.Offset,
		Limit:    a.Limit,
		AuthorID: a.AuthorID,
		Title:    a.Title,
	}
}
