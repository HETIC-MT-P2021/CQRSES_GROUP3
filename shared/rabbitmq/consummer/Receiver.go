package consummer

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

type ConsumeQueueService struct {
	Queue string
}

func (qs *ConsumeQueueService) ReceiveFromRabbit() {
	var messages []byte
	var messageList []byte

	url := os.Getenv("AMQP_URL")

	if url == "" {
		url = "amqp://user:bitnami@rabbitmq:5672"
	}
	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)
	fmt.Println(url)

	if err != nil {
		fmt.Println("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()

	if err != nil {
		fmt.Println("could not open RabbitMQ channel:" + err.Error())
	}

	msgs, err := channel.Consume(qs.Queue, "", false, false, false, false, nil)

	if err != nil {
		fmt.Println("error consuming the queue: " + err.Error())
	}

	// The msgs will be a go channel, not an amqp channel
	for msg := range msgs {
		messageList = append(messages, msg.Body...)
		message := string(msg.Body)
		fmt.Println("message received: " + message)
		msg.Ack(false)
	}

	for message := range messageList {
		fmt.Println(message)
	}

	defer connection.Close()
}
