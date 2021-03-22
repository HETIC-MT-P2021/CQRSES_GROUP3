package repositories

import (
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/services"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/es"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

// PersistArticle saves an article in es
func PersistArticle(article *models.Article) error {
	document := services.Document{
		Body: &es.Event{
			AggregateID: uuid.NewV4().String(),
			Typology:    es.Create,
			Payload:     article,
			CreatedAt:   time.Now(),
			Index:       1, // First event for this article so the index should be 1
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

func UpdateArticle(aggregateId string, article *models.Article) error {
	document := services.Document{
		Body: &es.Event{
			AggregateID: aggregateId,
			Typology:    es.Put,
			Payload:     article,
			CreatedAt:   time.Now(),
			Index:       1, // First event for this article so the index should be 1
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

func GetArticleEventByAggregateId(id string) ([]*es.Event, error) {
	//var article models.Article
	//var articleList []*models.Article
	var event es.Event
	var eventList []*es.Event
	query := services.ConstructBoolQuery("AggregateID", id)
	srl := services.SearchWithKeyword("article", query)
	for _, sr := range srl {
		payload := sr.Body.(map[string]interface{})
		err := helpers.Decode(payload, &event)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("cannot decode payload to Article struct, reason: %s", err.Error())
		}
		eventList = append(eventList, &event)
	}
	return eventList, nil
}
