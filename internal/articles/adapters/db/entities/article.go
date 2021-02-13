package entities

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/db/types"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain"
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

func (e *ArticleEntity) ToModel() *domain.Article {
	author := new(domain.Author)
	if e.Author != nil {
		author = e.Author.ToModel()
	}
	return &domain.Article{
		ID:        e.ID,
		Title:     e.Title,
		Content:   e.Content,
		Author:    author,
		IsDeleted: !e.DeletedAt.IsZero(),
	}
}

func NewArticleEntity(command *domain.ArticleCreateCommand) *ArticleEntity {
	return &ArticleEntity{
		Title:   command.Title,
		Content: command.Content,
	}
}

func (e *ArticleEntity) UpdateFrom(command *domain.ArticleUpdateCommand) {
	e.Title = command.Title
	e.Content = command.Content
}
