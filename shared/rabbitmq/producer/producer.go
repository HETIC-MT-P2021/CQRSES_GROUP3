package producer

import (
	"encoding/json"
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/rabbitmq"
	"os"
)

type QueueService struct {
	Queue string
	Data  interface{}
}

func (qs *QueueService) NewSendToRabbit() error {
	url := os.Getenv("AMQP_URL")

	if url == "" {
		url = "amqp://user:bitnami@rabbitmq:5672"
	}
	cfg := rabbitmq.Config{
		URL:        url,
		QueueName:  qs.Queue,
		Exchange:   "event",
		RoutingKey: "#",
	}
	queue, err := rabbitmq.NewQueueInstance(cfg)
	if err != nil {
		fmt.Println("error declaring the new queue instance: " + err.Error())
		return err
	}
	fmt.Println(qs.Data)
	data, err := json.Marshal(qs.Data)
	if err != nil {
		fmt.Println("error in json conversion of event: " + err.Error())
		return err
	}
	err = queue.Publish(data)
	if err != nil {
		fmt.Println("error publishing a message to the queue:" + err.Error())
		return err
	}
	return nil
}
