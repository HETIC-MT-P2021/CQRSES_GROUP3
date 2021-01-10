package controllers

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestSearchService(c *gin.Context) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "jack",
			},
		},
	}
	searchResults := services.SearchWithKeyword("articles", &query)
	c.JSON(http.StatusOK, searchResults)
	return
}
