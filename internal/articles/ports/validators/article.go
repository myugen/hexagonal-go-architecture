package validators

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain"
)

type ArticleValidator interface {
	ValidateCreate(command *domain.ArticleCreateCommand) error
	ValidateUpdate(command *domain.ArticleUpdateCommand) error
	ValidateDelete(model *domain.Article) error
	ValidateRecover(model *domain.Article) error
}
