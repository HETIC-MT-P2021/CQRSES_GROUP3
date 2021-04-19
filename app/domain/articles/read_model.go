package articles

import (
	"errors"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/es"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	"sort"
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

func (r *ReadModel) ProjectNewReadModel(eventList []*es.Event) (models.Article, error, int) {
	var articleStruct models.Article

	if len(eventList) < 1 {
		return models.Article{}, errors.New("event list is empty"), 0
	}
	// Sort slice by CreatedAt time.
	sort.SliceStable(eventList, func(i, j int) bool {
		return eventList[i].CreatedAt.Before(eventList[j].CreatedAt)
	})

	// Sort slice by Index.
	sort.SliceStable(eventList, func(i, j int) bool {
		return eventList[i].Index < eventList[j].Index
	})

	readModel := models.Article{
		AuthorID:  0,
		Title:     "",
		Content:   "",
		CreatedAt: time.Now(),
	}

	lastPlayedIndex := 0

	for _, event := range eventList {
		article := event.Payload.(map[string]interface{})
		err := helpers.Decode(article, &articleStruct)
		if err != nil {
			return models.Article{}, err, 0
		}
		lastPlayedIndex = int(event.Index)
		// if delete event return nil article
		if string(event.Typology) == "delete" {
			return models.Article{}, nil, 0
		}

		applyChanges(&readModel, &articleStruct)
	}

	return readModel, nil, lastPlayedIndex
}

func applyChanges(old *models.Article, new *models.Article) {
	// check each field
	if old.AuthorID != new.AuthorID {
		old.AuthorID = new.AuthorID
	}
	if old.Title != new.Title {
		old.Title = new.Title
	}
	if old.Content != new.Content {
		old.Content = new.Content
	}
	if old.CreatedAt != new.CreatedAt {
		old.CreatedAt = new.CreatedAt
	}
}
