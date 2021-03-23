package models

import (
	"errors"
	"time"
)

// Article model.
type Article struct {
	AuthorID    uint64
	Title       string
	Content     string
	CreatedAt   time.Time
	AggregateID string
}

// ArticleForm represents the informations needed to create an article
type ArticleForm struct {
	AuthorID uint64
	Title    string
	Content  string
}

// ValidateArticle takes an article form as parameter and check if its properties are valid
func ValidateArticle(article *ArticleForm) error {
	if article.Title == "" {
		return errors.New("Invalid title")
	} else if article.Content == "" {
		return errors.New("Invalid content")
	}

	return nil
}
