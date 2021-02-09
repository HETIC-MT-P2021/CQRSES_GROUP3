package rabbitmq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RabbitmqTestSuite struct {
	suite.Suite
	queue *Rabbitmq
}

func (s *RabbitmqTestSuite) SetupTest() {
	q, err := NewQueueInstance(Config{
		URL:        "amqp://user:bitnami@localhost:5672",
		Exchange:   "test_exchange",
		QueueName:  "test_queuename",
		RoutingKey: "test",
		BindingKey: "test",
	})
	if err != nil {
		panic(err)
	}
	s.queue = q
}

func (s *RabbitmqTestSuite) TearDownTest() {
	s.queue.Close()
}

func (s *RabbitmqTestSuite) TestPublishMessage() {
	s.T().Run("publish a message", func(t *testing.T) {
		message := []byte("Test")
		err := s.queue.Publish(message)
		assert.NoError(t, err, "Publish() error:\nwant  nil\ngot  %v", err)
	})
}

func (s *RabbitmqTestSuite) TestConsumeMessage() {
	message := []byte("Test")
	err := s.queue.Publish(message)
	if err != nil {
		panic(err)
	}
	messages, err := s.queue.Consume()
	if err != nil {
		panic(err)
	}

	s.T().Run("consume a message", func(t *testing.T) {
		assert.NoError(t, err, "Consume() error:\nwant  nil\ngot  %v", err)
	})

	s.T().Run("expect a delivery", func(t *testing.T) {
		expected := []byte("Test")
		select {
		case message := <-messages:
			{
				assert.Equal(t, expected, message, "Consume() error:\nwant  %v\ngot  %v", expected, message)
			}
		}
	})
}

func TestRabbitmqTestSuite(t *testing.T) {
	suite.Run(t, new(RabbitmqTestSuite))
}
