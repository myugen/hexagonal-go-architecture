package entities

import (
	"github.com/myugen/hexagonal-go-architecture/pkg/articles/adapters/db/types"
	"github.com/myugen/hexagonal-go-architecture/pkg/articles/domain/models"
)

type Article struct {
	tableName struct{} `pg:"alias:article"`

	ID      uint
	Title   string
	Content string
	Author  *Author `pg:"rel:has-one"`
	types.Datetime
}

func (e *Article) ToModel() *models.Article {
	return &models.Article{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		Author:  *e.Author.ToModel(),
	}
}
