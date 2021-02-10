package main

import (
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/rabbitmq/consummer"
)

func main() {
	resultsChannel := make(chan interface{})
	cqs := consummer.ConsumeQueueService{Queue: "CreateArticleCommand"}
	go cqs.Consume(resultsChannel)
	for data := range resultsChannel {
		fmt.Println(data)
	}
}
