package articles

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/repositories"
)

type GetArticleByAggregateIDQuery struct {
	AggregateID string
}

type ArticleQueryHandler struct{}

func (aqh *ArticleQueryHandler) Handle(command cqrs.QueryMessage) (interface{}, error, int) {
	switch cmd := command.Payload().(type) {
	case *GetArticleByAggregateIDQuery:
		r := ReadModel{AggregateID: cmd.AggregateID}
		eventList, err := repositories.GetArticleEventByAggregateId(r.AggregateID)
		if err != nil {
			return nil, err, 0
		}
		readModel, err, lastIndex := r.ProjectNewReadModel(eventList)
		return readModel, err, lastIndex
	default:
		return nil, nil, 0
	}
}

func NewArticleQueryHandler() *ArticleQueryHandler {
	return &ArticleQueryHandler{}
}
