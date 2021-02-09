package controllers

import (
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func TestSearchService(c *gin.Context) {
	// How to build a query
	id := c.Param("id")
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"term": map[string]interface{}{
						"AggregateID": id,
					},
				},
			},
		},
	}
	log.Info(query)
	searchResults := services.SearchWithKeyword("article", &query)
	if searchResults == nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("FUCK THIS SHIT"))
		return
	}
	c.JSON(http.StatusOK, searchResults)
	return
}
