package entities

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain"
)

type AuthorEntity struct {
	tableName struct{} `pg:"authors,alias:author"`

	ID   uint
	Name string
}

func (e *AuthorEntity) ToModel() *domain.Author {
	return &domain.Author{
		ID:   e.ID,
		Name: e.Name,
	}
}
