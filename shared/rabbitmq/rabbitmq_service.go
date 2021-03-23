package rabbitmq

import (
	"github.com/streadway/amqp"
)

// Config parameters for queue service
type Config struct {
	URL        string
	Exchange   string
	QueueName  string
	RoutingKey string
	BindingKey string
}

// Rabbitmq implements the Queue interface
type Rabbitmq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	config  Config
}

// Service interfaces the package methods
type Service interface {
	Publish(message []byte) error
	Consume() (<-chan []byte, error)
	Close() error
}

// NewQueueInstance set up the necessary queue instance to pass messages to the broker.
func NewQueueInstance(config Config) (*Rabbitmq, error) {
	conn, err := amqp.Dial(config.URL)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if _, err := NewQueue(ch, config.QueueName); err != nil {
		return nil, err
	}

	return &Rabbitmq{
		conn:    conn,
		channel: ch,
		config:  config,
	}, nil
}

// NewQueueInstanceWithBinding set up all the necessary instances to pass messages to the broker.
func NewQueueInstanceWithBinding(config Config) (*Rabbitmq, error) {
	conn, err := amqp.Dial(config.URL)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	if err := NewExchange(ch, config.Exchange); err != nil {
		return nil, err
	}

	if err := NewQueueWithBinding(ch, config.Exchange, config.QueueName, config.BindingKey); err != nil {
		return nil, err
	}

	return &Rabbitmq{
		conn:    conn,
		channel: ch,
		config:  config,
	}, nil
}

// NewExchange declare a new rabbit Exchange.
func NewExchange(c *amqp.Channel, name string) error {
	return c.ExchangeDeclare(
		name,    // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
}

// NewQueue declare a new rabbit Queue.
func NewQueue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return q, err
	}

	return q, nil
}

// NewQueueWithBinding declare a new rabbit Queue and bing it to the Exchange channel.
func NewQueueWithBinding(ch *amqp.Channel, exchange string, name string, bindingKey string) error {
	q, err := NewQueue(ch, name)
	if err != nil {
		return err
	}
	return ch.QueueBind(
		q.Name,
		bindingKey,
		exchange,
		false,
		nil,
	)
}

// Publish message to rabbitmq
func (r *Rabbitmq) Publish(message []byte) error {
	return r.channel.Publish(
		r.config.Exchange,
		r.config.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
}

// Consume message from rabbitmq
func (r *Rabbitmq) Consume() (<-chan []byte, error) {
	msgs, err := r.channel.Consume(
		r.config.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	deliveries := make(chan []byte)
	go func() {
		for msg := range msgs {
			deliveries <- msg.Body
		}
	}()
	return deliveries, nil
}

// Close rabbitmq connection
func (r *Rabbitmq) Close() error {
	if err := r.conn.Close(); err != nil {
		return err
	}
	if err := r.channel.Close(); err != nil {
		return err
	}
	return nil
}
