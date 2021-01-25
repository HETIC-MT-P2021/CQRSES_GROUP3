package articles

import "github.com/HETIC-MT-P2021/CQRSES_GROUP3/core/cqrs"

type GetArticleQuery struct {}

type CreateArticleQueryHandler struct {}

func (aqh *CreateArticleQueryHandler) Handle (command cqrs.QueryMessage) (interface{}, error) {
	switch command.Payload() {
	case GetArticleQuery{}:
		return nil, nil
	default:
		return nil, nil
	}
}

func NewArticleQueryHandler() *CreateArticleQueryHandler {
	return &CreateArticleQueryHandler{}
}
