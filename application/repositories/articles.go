package repositories

import (
	log "github.com/sirupsen/logrus"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/services"
	"errors"
	"time"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/core/es"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/models"
)

// PersistArticle saves an article in es
func PersistArticle(articleForm *models.ArticleForm) error {
	// Validate the form struct
	err := models.ValidateArticle(articleForm)
	if err != nil {
		return errors.New("error during article form validations")
	}
	
	// Create the final article, event and document
	article := models.Article {
		AuthorID: articleForm.AuthorID,
		Title: articleForm.Title,
		Content: articleForm.Content,
		CreatedAt: time.Now(),
	}
	document := services.Document {
		Body: es.Event {
			Typology: es.Create,
			Payload: article,
			CreatedAt: time.Now(),
			Index: 1,					// First event for this article so the index should be 1
		},
	}
		
	err = services.CreateNewDocumentInIndex("article", &document)
	if err != nil {
		log.Error("Error while creating event : ", err)
		return err
	}
	log.Info("document created : ", document)
	return nil
}

// Create Article 
// -> command create article (payload) 
// -> domain create article handler 
// -> decode payload 
// -> create event
// -> persist event