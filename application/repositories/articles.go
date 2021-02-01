package repositories

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/services"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/core/es"
	log "github.com/sirupsen/logrus"
	"time"
)

// PersistArticle saves an article in es
func PersistArticle(article *models.Article) error {
	document := services.Document{
		Body: es.Event{
			Typology:  es.Create,
			Payload:   article,
			CreatedAt: time.Now(),
			Index:     1, // First event for this article so the index should be 1
		},
	}

	err := services.CreateNewDocumentInIndex("article", &document)
	if err != nil {
		log.Error("Error while creating event : ", err)
		return err
	}
	log.Info("document created : ", document)
	return nil
}
