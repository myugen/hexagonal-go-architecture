package responses

import "github.com/myugen/hexagonal-go-architecture/pkg/articles/domain/models"

type Article struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func FromModel(article *models.Article) *Article {
	return &Article{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Author:  article.Author.Name,
	}
}
