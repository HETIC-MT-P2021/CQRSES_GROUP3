package articles

import "github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/cqrs"

type GetArticleQuery struct{}

type ArticleQueryHandler struct{}

func (aqh *ArticleQueryHandler) Handle(command cqrs.QueryMessage) (interface{}, error) {
	switch command.Payload() {
	case GetArticleQuery{}:
		return nil, nil
	default:
		return nil, nil
	}
}

func NewArticleQueryHandler() *ArticleQueryHandler {
	return &ArticleQueryHandler{}
}
