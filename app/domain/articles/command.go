package articles

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
)

// CreateArticleCommand is the struct we use to create a new command
type CreateArticleCommand struct {
	ArticleForm models.ArticleForm
}

type EditArticleCommand struct {
	AggregateId string
	ArticleForm models.ArticleForm
}
type DeleteArticleCommand struct{}

type ArticleCommandHandler struct{}

func (ach *ArticleCommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *CreateArticleCommand:
		article, err := validateAndPublishArticleEvent(&cmd.ArticleForm)
		if err != nil {
			return nil, err
		}
		return article, err 
	case *EditArticleCommand:
		article, err := validateAndPublishArticleVersion(cmd.AggregateId, &cmd.ArticleForm)
		if err != nil {
			return nil, err
		}
		return article, err
	default:
		return nil, nil
	}
}

func NewArticleCommandHandler() *ArticleCommandHandler {
	return &ArticleCommandHandler{}
}
