package articles

import "github.com/HETIC-MT-P2021/CQRSES_GROUP3/core/cqrs"

type CreateArticleCommand struct {}
type EditArticleCommand struct {}
type DeleteArticleCommand struct {}

type ArticleCommandHandler struct {}

func (ach *ArticleCommandHandler) Handle (command cqrs.CommandMessage) (interface{}, error) {
	switch command.Payload() {
	case CreateArticleCommand{}:
		return nil, nil
	case EditArticleCommand{}:
		return nil, nil
	default:
		return nil, nil
	}
}

func NewArticleCommandHandler() *ArticleCommandHandler {
	return &ArticleCommandHandler{}
}