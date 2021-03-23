package articles

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/es"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/rabbitmq/producer"
	"github.com/satori/go.uuid"
	"time"
)

func validateAndPublishArticleEvent(articleForm *models.ArticleForm) (models.Article, error) {
	if err := models.ValidateArticle(articleForm); err != nil {
		return models.Article{}, err
	}

	aggregateID := uuid.NewV4().String()

	article := models.Article{
		AggregateID: aggregateID,
		AuthorID:    articleForm.AuthorID,
		Title:       articleForm.Title,
		Content:     articleForm.Content,
		CreatedAt:   time.Now(),
	}

	event := es.Event{
		AggregateID: aggregateID,
		Typology:    es.Create,
		Payload:     article,
		CreatedAt:   time.Now(),
		Index:       1, // First event for this article so the index should be 1
	}

	queue := producer.QueueService{
		Queue: string(event.Typology),
		Data:  event,
	}
	err := queue.NewSendToRabbit()
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func validateAndPublishArticleVersion(aggregateId string, articleForm *models.ArticleForm) (models.Article, error) {
	if err := models.ValidateArticle(articleForm); err != nil {
		return models.Article{}, err
	}

	article := models.Article{
		AuthorID:    articleForm.AuthorID,
		Title:       articleForm.Title,
		Content:     articleForm.Content,
		CreatedAt:   time.Now(),
		AggregateID: aggregateId,
	}

	event := es.Event{
		AggregateID: aggregateId,
		Typology:    es.Put,
		Payload:     article,
		CreatedAt:   time.Now(),
		Index:       1, // First event for this article so the index should be 1
	}

	queue := producer.QueueService{
		Queue: string(event.Typology),
		Data:  event,
	}
	err := queue.NewSendToRabbit()
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func publishDeleteArticleEvent(aggregateID string) error {
	event := es.Event{
		AggregateID: aggregateID,
		Typology:    es.Delete,
		Payload:     models.Article{},
		CreatedAt:   time.Now(),
		Index:       1,
	}

	queue := producer.QueueService{
		Queue: string(event.Typology),
		Data:  event,
	}

	return queue.NewSendToRabbit()
}
