package producer

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type QueueService struct {
	Queue string
	Data interface{}
}

func (qs *QueueService) SendToRabbit() error{

	url := "amqp://user:bitnami@rabbitmq:5672"

	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)

	if err != nil {
		fmt.Println("could not establish connection with RabbitMQ:" + err.Error())
		return err
	}

	// Create a channel from the connection. We'll use channels to access the data in the queue rather than the
	channel, err := connection.Channel()

	if err != nil {
		fmt.Println("could not open RabbitMQ channel:" + err.Error())
		return err
	}

	// We create an exchange that will bind to the queue to send and receive messages
	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	if err != nil {
		fmt.Println("could not declare exchange:" + err.Error())
		return err
	}

	// We create a queue named Campaigns
	_, err = channel.QueueDeclare(qs.Queue, true, false, false, false, nil)

	if err != nil {
		fmt.Println("error declaring the queue: " + err.Error())
		return err
	}

	// We create a message to be sent to the queue.
	// It has to be an instance of the aqmp publishing struct
	data, err := json.Marshal(qs.Data)
	if err != nil{
		fmt.Println("error in json conversion of event: " + err.Error())
		return err
	}
	message := amqp.Publishing{
		Body: data,
	}
	// We publish the message to the exahange we created earlier
	err = channel.Publish("events", "random-key", false, false, message)

	if err != nil {
		fmt.Println("error publishing a message to the queue:" + err.Error())
		return err
	}

	//
	err = channel.QueueBind(qs.Queue, "#", "events", false, nil)

	if err != nil {
		fmt.Println("error binding to the queue: " + err.Error())
		return err
	}

	return nil
}
