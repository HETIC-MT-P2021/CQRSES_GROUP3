package articles

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/domain/articles"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/cqrs"
)

type GetArticleByAggregateIDQuery struct {
	AggregateID string
}

type ArticleQueryHandler struct{}

func (aqh *ArticleQueryHandler) Handle(command cqrs.QueryMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *GetArticleByAggregateIDQuery:
<<<<<<< HEAD:app/domain/articles/query.go
		r := articles.ReadModel{AggregateID: cmd.AggregateID}
		r.ProjectNewReadModel()
		return nil, nil
=======
		r := ReadModel{AggregateID: cmd.AggregateID}
		readModel, err := r.ProjectNewReadModel()
		return readModel, err
>>>>>>> feat: finished read model construction:domain/articles/query.go
	default:
		return nil, nil
	}
}

func NewArticleQueryHandler() *ArticleQueryHandler {
	return &ArticleQueryHandler{}
}
