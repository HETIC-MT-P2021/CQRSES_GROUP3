package models

import (
	"errors"
	"time"
)

type Article struct {
	AuthorID 	uint64
	Title 		string
	Content		string
	CreatedAt	time.Time
}

type ArticleForm struct {
	AuthorID 	uint64
	Title 		string
	Content 	string
}

// ValidateArticle takes an article form as parameter and check if its properties are valid
func ValidateArticle(article *ArticleForm) error {
	if article.Title == "" {
		return errors.New("Invalid title")
	} elif article.Content == "" {
		return errors.New("Invalid content")
	}

	return nil
}