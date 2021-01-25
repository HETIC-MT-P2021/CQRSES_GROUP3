package repositories

import (
	"errors"
	"time"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/core/es"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/models"
)

// PersistArticle saves an article in es
func PersistArticle(articleForm *models.ArticleForm) error {
	err := models.ValidateArticle(articleForm)
	if err != nil {
		return errors.New("error during article form validations")
	}
	
	article := models.Article {
		AuthorID: articleForm.AuthorID,
		Title: articleForm.Title,
		Content: articleForm.Content,
		CreatedAt: time.Now(),
	}
	event := es.Event {
		Typology: es.Create,
		Payload: article,
		CreatedAt: time.Now(),
	}

	return nil
}

// Create Article 
// -> command create article (payload) 
// -> domain create article handler 
// -> decode payload 
// -> create event
// -> persist event