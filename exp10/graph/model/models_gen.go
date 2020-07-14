// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ArticleID string `json:"articleId"`
	Content   string `json:"content"`
}

type NewArticle struct {
	ArticleID string `json:"articleId"`
	Content   string `json:"content"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}