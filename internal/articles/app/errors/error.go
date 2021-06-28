package errors

import (
	"fmt"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/errors"
)

type ArticleError struct {
	code    ArticleErrorCode
	message string
	cause   error
}

func (e *ArticleError) Code() ArticleErrorCode {
	return e.code
}

func (e *ArticleError) Message() string {
	return e.message
}

func (e *ArticleError) WithMessage(message string) *ArticleError {
	e.message = message
	return e
}

func (e *ArticleError) Cause() error {
	return e.cause
}

func (e *ArticleError) WithCause(err error) *ArticleError {
	e.cause = err
	return e
}

func (e *ArticleError) Unwrap() error {
	return e.Cause()
}

func (e *ArticleError) Error() string {
	errorMsg := e.code.String()
	if e.message != "" {
		errorMsg = fmt.Sprintf(errorMsg+": %s", e.message)
	}

	return errorMsg
}

func (e *ArticleError) Context() map[string]interface{} {
	ctx := map[string]interface{}{
		"code":    e.code,
		"message": e.message,
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

func NewArticleError(code ArticleErrorCode, message string, cause ...error) *ArticleError {
	if len(cause) > 0 {
		return &ArticleError{code: code, message: message, cause: cause[0]}
	}
	return &ArticleError{code: code, message: message}
}
