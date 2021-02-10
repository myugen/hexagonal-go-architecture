package entities

import "github.com/myugen/hexagonal-go-architecture/internal/articles/domain/models"

type AuthorEntity struct {
	tableName struct{} `pg:"authors,alias:author"`

	ID   uint
	Name string
}

func (e *AuthorEntity) ToModel() *models.Author {
	return &models.Author{
		ID:   e.ID,
		Name: e.Name,
	}
}
