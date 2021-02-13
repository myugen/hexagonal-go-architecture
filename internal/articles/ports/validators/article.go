package validators

import "github.com/myugen/hexagonal-go-architecture/internal/articles/domain/models"

type ArticleValidator interface {
	ValidateCreate(command *models.ArticleCreateCommand) error
	ValidateUpdate(command *models.ArticleUpdateCommand) error
	ValidateDelete(model *models.Article) error
	ValidateRecover(model *models.Article) error
}
