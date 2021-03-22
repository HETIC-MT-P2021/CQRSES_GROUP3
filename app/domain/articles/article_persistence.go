package articles

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/repositories"
	"time"
)

func validateAndPersistArticle(articleForm *models.ArticleForm) (models.Article, error) {
	if err := models.ValidateArticle(articleForm); err != nil {
		return models.Article{}, err
	}

	article := models.Article{
		AuthorID:  articleForm.AuthorID,
		Title:     articleForm.Title,
		Content:   articleForm.Content,
		CreatedAt: time.Now(),
	}

	if err := repositories.PersistArticle(&article); err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func validateAndUpdateArticle(aggregateId string, articleForm *models.ArticleForm) (models.Article, error) {
	if err := models.ValidateArticle(articleForm); err != nil {
		return models.Article{}, err
	}

	article := models.Article{
		AuthorID:  articleForm.AuthorID,
		Title:     articleForm.Title,
		Content:   articleForm.Content,
		CreatedAt: time.Now(),
	}

	if err := repositories.UpdateArticle(aggregateId, &article); err != nil {
		return models.Article{}, err
	}

	return article, nil
}
