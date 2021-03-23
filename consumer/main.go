package main

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/es"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/rabbitmq/consummer"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/repositories"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Consuming queues")

	// DB config
	ecfg := database.EsCfg{Url: "http://es:9200"}
	database.GetESClient(&ecfg)

	// Consumer creation
	var queues []es.Typology
	queues = append(queues, es.Create)
	queues = append(queues, es.Put)

	createConsumersByEventTypologies(queues)
}

// createConsumersByEventTypologies create queues instance with name based on the event typologies
// uses the channel of result to decode the data and persist them in es.
func createConsumersByEventTypologies(typologies []es.Typology)  {
	resultsChannel := make(chan interface{})
	for _, typology := range typologies{
		queue := consummer.ConsumeQueueService{Queue: string(typology)}
		go queue.Consume(resultsChannel)
	}
	for data := range resultsChannel {
		decodeAndPersistArticle(data)
	}
}

// decodeAndPersistArticle takes an interface and decode it to an event.
// uses this event to create a new document in es.
func decodeAndPersistArticle(data interface{})  {
	var eventStruct es.Event
	err := helpers.Decode(data, &eventStruct)
	if err != nil {
		log.Error("Error while decoding event: ", err)
	}
	switch eventStruct.Typology {
	case es.Create:
		err = repositories.PersistArticleEvent(&eventStruct)
		if err != nil {
			log.Error("Error while persisting event: ", err)
		}
	case es.Put:
		err = repositories.PersistArticleVersionEvent(eventStruct.AggregateID,&eventStruct)
		if err != nil {
			log.Error("Error while persisting versioning event: ", err)
		}
	}

}