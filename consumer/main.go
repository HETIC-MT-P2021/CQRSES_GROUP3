package main

import (
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/rabbitmq/consummer"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Consuming queues")
	resultsChannel := make(chan interface{})
	cqs := consummer.ConsumeQueueService{Queue: "create"}
	go cqs.Consume(resultsChannel)
	for data := range resultsChannel {
		fmt.Println(data)
	}
}
