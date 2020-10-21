package services

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrExist = errors.New("Article already exist")
)

type IArticle interface {
	Get(ctx context.Context)
}
