package main

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/rabbitmq/consummer"
)

func main() {
	cqs := consummer.ConsumeQueueService{Queue: "CreateArticleCommand"}
	cqs.Consume()
}
