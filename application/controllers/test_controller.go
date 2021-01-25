package controllers

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestSearchService(c *gin.Context) {
	// How to build a query
	_ = map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "jack",
			},
		},
	}
	doc, err := services.GetDocumentById("articles", "Cuw3EncB4DCj_NoRaliY")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, doc)
	return
}
