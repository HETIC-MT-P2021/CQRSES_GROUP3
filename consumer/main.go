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
	ecfg := database.EsCfg{Url: "http://localhost:9200"}
	database.GetESClient(&ecfg)
	resultsChannel := make(chan interface{})
	cqs := consummer.ConsumeQueueService{Queue: string(es.Put)}
	go cqs.Consume(resultsChannel)
	for data := range resultsChannel {
		decodeAndPersistArticle(data)
	}
}

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