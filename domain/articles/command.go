package articles

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/core/cqrs"
)

// CreateArticleCommand is the struct we use to create a new command
type CreateArticleCommand struct {
	ArticleForm models.ArticleForm
}

func (c CreateArticleCommand) CommandType() string {
	return "CreateArticleCommand"
}
func (c CreateArticleCommand) Payload() interface{} {
	return &c
}

type EditArticleCommand struct {}
type DeleteArticleCommand struct {}

type ArticleCommandHandler struct {}

func (ach *ArticleCommandHandler) Handle (command cqrs.CommandMessage) (interface{}, error) {
	switch command.CommandType() {
	case "CreateArticleCommand":
		payload := command.Payload().(*CreateArticleCommand)
		article, err := validateAndPersistArticle(&payload.ArticleForm)
		return article, err
	case "EditArticleCommand":
		return nil, nil
	default:
		return nil, nil
	}
}

func NewArticleCommandHandler() *ArticleCommandHandler {
	return &ArticleCommandHandler{}
}