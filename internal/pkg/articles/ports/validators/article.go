package validators

import "github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/domain/models"

type IArticle interface {
	ValidateCreate(command *models.ArticleCreateCommand) error
	ValidateUpdate(command *models.ArticleUpdateCommand) error
	ValidateDelete(model *models.Article) error
	ValidateRecover(model *models.Article) error
}
