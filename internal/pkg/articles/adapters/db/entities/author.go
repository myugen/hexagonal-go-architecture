package entities

import "github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"

type Author struct {
	tableName struct{} `pg:"alias:author"`

	ID   uint
	Name string
}

func (e *Author) ToModel() *models.Author {
	return &models.Author{
		ID:   e.ID,
		Name: e.Name,
	}
}
