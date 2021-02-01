package articles

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/core/cqrs"
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
		article, err := validateAndPersistArticle(&cmd.ArticleForm)
		return article, err
	case *EditArticleCommand:
		return nil, nil
	default:
		return nil, nil
	}
}

func NewArticleCommandHandler() *ArticleCommandHandler {
	return &ArticleCommandHandler{}
}
