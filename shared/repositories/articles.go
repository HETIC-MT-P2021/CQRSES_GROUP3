package repositories

import (
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/services"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/es"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/helpers"
	log "github.com/sirupsen/logrus"
)

// PersistArticleEvent saves an es.Event in es containing the article as a payload
func PersistArticleEvent(event *es.Event) error {
	document := services.Document{
		Body: event,
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
	var eventList []*es.Event
	query := services.ConstructBoolQuery("AggregateID", id)
	srl := services.SearchWithKeyword("article", query)
	for _, sr := range srl {
		var event es.Event
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
