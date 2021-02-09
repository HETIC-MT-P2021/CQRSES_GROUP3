package controllers

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/domain/articles"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/core/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
