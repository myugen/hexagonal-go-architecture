package types

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"
)

type ArticleFilter struct {
	Limit           int
	Offset          int
	AuthorID        uint
	Title           string
	IncludedDeleted bool
}

func NewArticleFilter(query *models.ArticleQuery) *ArticleFilter {
	return &ArticleFilter{
		Limit:           query.Limit,
		Offset:          query.Offset * query.Limit,
		AuthorID:        query.AuthorID,
		Title:           query.Title,
		IncludedDeleted: query.IncludedDeleted,
	}
}

func (a *ArticleFilter) Where(q *orm.Query) (*orm.Query, error) {
	if a.AuthorID > 0 {
		q.Where("author_id = ?", a.AuthorID)
	}

	if a.Title != "" {
		q.Where("title ILIKE ?", "%"+a.Title+"%")
	}

	if a.Limit > 0 {
		q.Limit(a.Limit).Offset(a.Offset)
	}

	if a.IncludedDeleted {
		q.AllWithDeleted()
	}

	return q, nil
}
