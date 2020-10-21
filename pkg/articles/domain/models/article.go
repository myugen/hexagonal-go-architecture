package models

type Article struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

type ArticleBuilder struct {
	Article
}

func (b *ArticleBuilder) Title(title string) *ArticleBuilder {
	b.Article.Title = title
	return b
}

func (b *ArticleBuilder) Content(content string) *ArticleBuilder {
	b.Article.Content = content
	return b
}

func (b *ArticleBuilder) Author(author string) *ArticleBuilder {
	b.Article.Author.Name = author
	return b
}

func (b *ArticleBuilder) Build() Article {
	return b.Article
}
