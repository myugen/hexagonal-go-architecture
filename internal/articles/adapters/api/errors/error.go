package errors

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/errors"
)

type ArticleAPIError struct {
	*echo.HTTPError
	cause error
}

func (e *ArticleAPIError) Cause() error {
	return e.cause
}

func (e *ArticleAPIError) Unwrap() error {
	return e.cause
}

func (e *ArticleAPIError) Context() map[string]interface{} {
	ctx := map[string]interface{}{
		"message": e.Message,
	}

	if causer := errors.Unwrap(e); causer != nil {
		if causerCtx := errors.Context(causer); len(causerCtx) > 0 {
			ctx["cause"] = causerCtx
		} else {
			ctx["cause"] = causer.Error()
		}
	}

	return ctx
}

func (e *ArticleAPIError) Error() string {
	return fmt.Sprintf("ArticleHanderError: %s", e.Message)
}

func NewArticleAPIError(httpstatus int, message string, cause ...error) *ArticleAPIError {
	if len(cause) > 0 {
		return &ArticleAPIError{HTTPError: echo.NewHTTPError(httpstatus, message), cause: cause[0]}
	}

	return &ArticleAPIError{HTTPError: echo.NewHTTPError(httpstatus, message)}
}
