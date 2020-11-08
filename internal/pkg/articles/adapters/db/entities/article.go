package entities

import (
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/db/types"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"
)

type ArticleEntity struct {
	tableName struct{} `pg:"articles,alias:article"`

	ID      uint
	Title   string
	Content string
	Author  *AuthorEntity `pg:"rel:has-one,fk:author"`
	types.Datetime
	types.SoftDelete
}

func (e *ArticleEntity) ToModel() *models.Article {
	author := new(models.Author)
	if e.Author != nil {
		author = e.Author.ToModel()
	}
	return &models.Article{
		ID:        e.ID,
		Title:     e.Title,
		Content:   e.Content,
		Author:    author,
		IsDeleted: !e.DeletedAt.IsZero(),
	}
}

func NewArticleEntity(command *models.ArticleCreateCommand) *ArticleEntity {
	return &ArticleEntity{
		Title:   command.Title,
		Content: command.Content,
	}
}

func (e *ArticleEntity) UpdateFrom(command *models.ArticleUpdateCommand) {
	e.Title = command.Title
	e.Content = command.Content
}
