package articles

import (
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/repositories"
	"time"
)

type ReadModel struct {
	AggregateID  string
	FinalPayload interface{}
	CreatedAt    time.Time
}

type ArticleReadModel struct {
	ReadModel ReadModel
}

func (r *ReadModel) ProjectNewReadModel() {
	var articleStruct models.Article
	eventList, err := repositories.GetArticleEventByAggregateId(r.AggregateID)
	if err != nil {
		fmt.Printf("Error while fetching article with Aggregate id: %s\n", r.AggregateID)
	}
	for _, event := range eventList {
		article := event.Payload.(map[string]interface{})
		err := helpers.Decode(article, &articleStruct)
		if err != nil {
			fmt.Println("ERROR", err)
		}
	}
}
