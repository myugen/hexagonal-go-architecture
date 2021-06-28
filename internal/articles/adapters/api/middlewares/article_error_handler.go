package middlewares

import (
	"net/http"

	articleError "github.com/myugen/hexagonal-go-architecture/internal/articles/app/errors"

	"github.com/labstack/echo/v4"
	articleAPIError "github.com/myugen/hexagonal-go-architecture/internal/articles/adapters/api/errors"
)

func ArticleErrorHandlerMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				if e, ok := err.(*articleError.ArticleError); ok {
					return apiErrorFromDomainError(e)
				} else {
					return echo.NewHTTPError(http.StatusInternalServerError, "Internal service error")
				}
			}
			return nil
		}
	}
}

func apiErrorFromDomainError(err *articleError.ArticleError) *articleAPIError.ArticleAPIError {
	type articleApiErrorResponse struct {
		Code    int
		Message string
	}
	var articleApiErrorResponseByErrorCode = map[articleError.ArticleErrorCode]articleApiErrorResponse{
		articleError.ArticleCreationErrorCode: {
			Code:    400,
			Message: "An error occurs creating an article",
		},
		articleError.ArticleUpdateErrorCode: {
			Code:    400,
			Message: "An error occurs updating an article",
		},
		articleError.ArticleDeletionErrorCode: {
			Code:    400,
			Message: "An error occurs deleting an article",
		},
		articleError.ArticleRecoveryErrorCode: {
			Code:    400,
			Message: "An error occurs recovering an article",
		},
		articleError.ArticleRetrievalErrorCode: {
			Code:    404,
			Message: "Article not found",
		},
	}

	errorResponse := articleApiErrorResponseByErrorCode[err.Code()]
	return articleAPIError.NewArticleAPIError(errorResponse.Code, errorResponse.Message, err)
}
