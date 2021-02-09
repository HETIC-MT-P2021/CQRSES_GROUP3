package consummer

import (
	"encoding/json"
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/rabbitmq"
	"os"
)

func (cqs *ConsumeQueueService) Consume() {
	url := os.Getenv("AMQP_URL")

	if url == "" {
		url = "amqp://user:bitnami@localhost:5672"
	}
	cfg := rabbitmq.Config{
		URL:        url,
		QueueName:  cqs.Queue,
		Exchange:   "event",
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
		fmt.Println(data)
	}
}
