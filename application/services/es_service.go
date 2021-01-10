package services

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/database"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	log "github.com/sirupsen/logrus"
)

type EsConnector struct {
}

type EsService interface {
	SearchWithKeyword(index string, query map[string]interface{})
}

type SearchResult struct {
	ID   interface{}
	Body interface{}
}

func SearchWithKeyword(index string, query *map[string]interface{}) *[]SearchResult {
	client, err := database.GetESClient()
	if err != nil {
		log.Error("Could not get elastic search client: %v", err)
		return nil
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Error("Error encoding query: %s", err)
		return nil
	}
	response, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(index),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil
	}
	prettifyNotFoundError(response)

	return mapSearchResults(response)
}

func prettifyNotFoundError(response *esapi.Response)  {
	if response.IsError() {
		var body map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
			log.Fatal("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatal("[%s] %s: %s",
				response.Status(),
				body["error"].(map[string]interface{})["type"],
				body["error"].(map[string]interface{})["reason"],
			)
		}
	}
}

func mapSearchResults(response *esapi.Response) *[]SearchResult {
	var body map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		log.Fatal("Error parsing the response body: %s", err)
		return nil
	}
	// Print the response status, number of results.
	log.Printf(
		"[%s] %d hits;",
		response.Status(),
		int(body["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	)
	var searchResultList []SearchResult
	// Print the ID and document source for each hit.
	for _, hit := range body["hits"].(map[string]interface{})["hits"].([]interface{}) {
		searchResult := SearchResult{
			ID:   hit.(map[string]interface{})["_id"],
			Body: hit.(map[string]interface{})["_source"],
		}
		searchResultList = append(searchResultList, searchResult)
	}

	return &searchResultList
}