package articles

import (
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/rabbitmq/producer"
)

// CreateArticleCommand is the struct we use to create a new command
type CreateArticleCommand struct {
	ArticleForm models.ArticleForm
}

type EditArticleCommand struct{}
type DeleteArticleCommand struct{}

type ArticleCommandHandler struct{}

func (ach *ArticleCommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *CreateArticleCommand:
		queue := producer.QueueService{
			Queue: command.CommandType(),
			Data: cmd.ArticleForm,
		}
		err := queue.SendToRabbit()
		fmt.Println(err)

		if err != nil {
			return nil, err
		}
		return nil, nil
	case *EditArticleCommand:
		return nil, nil
	default:
		return nil, nil
	}
}

func NewArticleCommandHandler() *ArticleCommandHandler {
	return &ArticleCommandHandler{}
}
