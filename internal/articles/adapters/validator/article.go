package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain"
	"github.com/pkg/errors"
)

var (
	errAlreadyDeleted = errors.New("article was already deleted")
	errNotDeleted     = errors.New("article is not deleted")
)

type ArticleValidator struct {
	*validator.Validate
}

func NewArticleValidator() *ArticleValidator {
	return &ArticleValidator{validator.New()}
}

func (v *ArticleValidator) ValidateCreate(command *domain.ArticleCreateCommand) error {
	return v.Struct(command)
}

func (v *ArticleValidator) ValidateUpdate(command *domain.ArticleUpdateCommand) error {
	return v.Struct(command)
}

func (v *ArticleValidator) ValidateDelete(model *domain.Article) error {
	if model.IsDeleted {
		return errAlreadyDeleted
	}
	return nil
}

func (v *ArticleValidator) ValidateRecover(model *domain.Article) error {
	if !model.IsDeleted {
		return errNotDeleted
	}
	return nil
}
