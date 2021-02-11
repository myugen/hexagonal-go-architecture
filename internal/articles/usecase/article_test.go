package usecase_test

import (
	"testing"

	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain/models"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/services"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/usecase"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestArticleUsecase(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(message string, _ ...int) {
		g.Fail(message)
	})

	var articleUsecase services.IArticle
	var ctx usecase.ArticleUsecaseContext

	g.Describe("Article use cases", func() {
		g.BeforeEach(func() {
			articleUsecase = usecase.NewArticleUsecase()
			ctx = usecase.NewArticleUsecaseContextMock()
		})

		g.Describe("Create method", func() {
			g.It("foo", func() {
				command := new(models.ArticleCreateCommand)
				result, err := articleUsecase.Create(ctx, command)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).ShouldNot(BeNil())
				Expect(result.ID).ShouldNot(BeEmpty())
			})
		})
	})
}
