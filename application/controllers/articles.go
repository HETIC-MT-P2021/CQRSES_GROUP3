package controllers

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/domain/articles"
	"net/http"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/models"
	"github.com/gin-gonic/gin"
)

// CreateArticle is the controller to handle the creation of an article
func CreateArticle(c *gin.Context) {
	var articleForm models.ArticleForm
	if err := c.ShouldBindJSON(&articleForm); err != nil {
		c.JSON(http.StatusBadRequest, "Missing fields for the article : " + err.Error())
		return
	}
	command := articles.CreateArticleCommand {
		ArticleForm: articleForm,
	}
	res, err := domain.Cb.Dispatch(command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}