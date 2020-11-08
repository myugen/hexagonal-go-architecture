package models

type Article struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

type ArticleQuery struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	AuthorID uint   `json:"author_id"`
	Title    string `json:"title"`
}
