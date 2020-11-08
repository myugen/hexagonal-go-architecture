package responses

import "github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"

type ArticleResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	IsDeleted bool   `json:"is_deleted"`
}

func NewArticleResponse(article *models.Article) *ArticleResponse {
	author := ""
	if article.Author != nil {
		author = article.Author.Name
	}
	return &ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Author:    author,
		IsDeleted: article.IsDeleted,
	}
}
