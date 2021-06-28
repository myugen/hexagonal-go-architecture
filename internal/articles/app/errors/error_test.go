package errors_test

import (
	"testing"

	"github.com/myugen/hexagonal-go-architecture/internal/articles/app/errors"

	"github.com/stretchr/testify/assert"
)

func TestArticleError_Context(t *testing.T) {
	assertt := assert.New(t)

	foo := errors.NewArticleError(errors.ArticleCreationErrorCode, "foo error")
	bar := errors.NewArticleError(errors.ArticleCreationErrorCode, "bar error", foo)
	baz := errors.NewArticleError(errors.ArticleCreationErrorCode, "bar error", bar)

	barExpectedCtx := bar.Context()
	bazCtx := baz.Context()

	assertt.Contains(bazCtx, "cause")
	assertt.Equal(bazCtx["cause"], barExpectedCtx)
}
