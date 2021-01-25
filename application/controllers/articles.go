package controllers

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/repositories"
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

	if err := repositories.PersistArticle(&articleForm); err != nil {
		c.JSON(http.StatusInternalServerError, "Error while saving article")
		return
	}

	c.JSON(http.StatusOK, "article created !")
}