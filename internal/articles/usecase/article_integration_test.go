// +build integration

package usecase_test

import (
	"testing"

	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain"

	"github.com/myugen/hexagonal-go-architecture/internal/articles/ports/services"
	"github.com/myugen/hexagonal-go-architecture/internal/articles/usecase"

	"github.com/franela/goblin"
)

func TestArticleUsecase(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(message string, _ ...int) {
		g.Fail(message)
	})

	var articleUsecase services.ArticleService
	var ctx services.ArticleServiceContext

	g.Describe("Article use cases", func() {
		g.BeforeEach(func() {
			articleUsecase = usecase.NewArticleUsecase()
			ctx = usecase.NewArticleUsecaseContextMock()
		})

		g.Describe("Create method", func() {
			g.It("foo", func() {
				command := domain.ArticleCreateCommand{
					Title:   "Foo",
					Content: "Lorem ipsum si amet",
				}
				result, err := articleUsecase.Create(ctx, &command)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).ShouldNot(BeNil())
				Expect(result.ID).ShouldNot(BeEmpty())
				Expect(result.Title).Should(Equal(command.Title))
				Expect(result.Content).Should(Equal(command.Content))
			})
		})
	})
}
