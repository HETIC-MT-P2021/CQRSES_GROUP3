package articles

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/cqrs"
)

type GetArticleByAggregateIDQuery struct {
	AggregateID string
}

type ArticleQueryHandler struct{}

func (aqh *ArticleQueryHandler) Handle(command cqrs.QueryMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *GetArticleByAggregateIDQuery:
		r := ReadModel{AggregateID: cmd.AggregateID}
		readModel, err := r.ProjectNewReadModel()
		return readModel, err
	default:
		return nil, nil
	}
}

func NewArticleQueryHandler() *ArticleQueryHandler {
	return &ArticleQueryHandler{}
}
