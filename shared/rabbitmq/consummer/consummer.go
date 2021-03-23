package consummer

import (
	"encoding/json"
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/rabbitmq"
	"os"
)

type ConsumeQueueService struct {
	Queue string
}

// Consume implements the ConsumeQueueService to get the targeted queue and takes a channel to retrieve the message
// Mostly used as a goroutine to create a concurrent function in order to listen on all messages received by the specified queue.
func (cqs *ConsumeQueueService) Consume(results chan<- interface{}) {
	url := os.Getenv("AMQP_URL")

	if url == "" {
		url = "amqp://user:bitnami@rabbitmq:5672"
	}
	cfg := rabbitmq.Config{
		URL:        url,
		QueueName:  cqs.Queue,
		Exchange:   "",
		RoutingKey: "#",
	}
	queue, err := rabbitmq.NewQueueInstance(cfg)
	if err != nil {
		fmt.Println("error declaring the new queue instance: " + err.Error())
	}
	messages, err := queue.Consume()
	if err != nil {
		fmt.Println("error consuming the queue: " + err.Error())
	}
	for message := range messages {
		var data interface{}
		if err := json.Unmarshal(message, &data); err != nil {
			fmt.Println("error in byte conversion of event: " + err.Error())
		}
		results <- data
	}
}
