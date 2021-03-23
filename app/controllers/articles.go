package controllers

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/domain/articles"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	"github.com/gin-gonic/gin"
	"net/http"
	log "github.com/sirupsen/logrus"
)

// CreateArticle is the controller to handle the creation of an article
func CreateArticle(c *gin.Context) {
	var articleForm models.ArticleForm
	if err := c.ShouldBindJSON(&articleForm); err != nil {
		c.JSON(http.StatusBadRequest, "Missing fields for the article : "+err.Error())
		return
	}

	command := articles.CreateArticleCommand{
		ArticleForm: articleForm,
	}
	cmdDescriptor := cqrs.NewCommandMessage(&command)
	res, err := domain.Cb.Dispatch(cmdDescriptor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateArticle
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var articleForm models.ArticleForm
	if err := c.ShouldBindJSON(&articleForm); err != nil {
		c.JSON(http.StatusBadRequest, "Missing fields for the article : "+err.Error())
		return
	}
	command := articles.EditArticleCommand{
		AggregateId: id,
		ArticleForm: articleForm,
	}
	cmdDescriptor := cqrs.NewCommandMessage(&command)
	article, err := domain.Cb.Dispatch(cmdDescriptor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, article)
	return
}

type ArticleResponse struct {
	Article   interface{}
	LastIndex int
}

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	query := articles.GetArticleByAggregateIDQuery{AggregateID: id}
	queryDescriptor := cqrs.NewQueryMessage(&query)
	article, err, lastIndex := domain.Qb.Dispatch(queryDescriptor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	emptyArticle := models.Article{}
	if article == emptyArticle {
		c.JSON(http.StatusNotFound, "Article does not exist")
		return
	}

	c.JSON(http.StatusOK, ArticleResponse{Article: article, LastIndex: lastIndex})
	return
}

func DeleteArticleById(c *gin.Context) {
	id := c.Param("id")
	command := articles.DeleteArticleCommand{AggregateID: id}
	cmdDescriptor := cqrs.NewCommandMessage(&command)

	_, err := domain.Cb.Dispatch(cmdDescriptor)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "Couldn't delete the article")
		return
	}

	c.JSON(http.StatusOK, "Article deleted")
}
