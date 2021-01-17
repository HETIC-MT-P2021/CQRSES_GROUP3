package controllers

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestSearchService(c *gin.Context) {
	searchResults := services.SearchWithKeyword("articles", "name", "jack", 10)
	c.JSON(http.StatusOK, searchResults)
	return
}
