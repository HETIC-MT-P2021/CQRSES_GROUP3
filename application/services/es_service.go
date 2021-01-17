package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/helpers"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type EsConnector struct {
}

// EsService interface.
type EsService interface {
	SearchWithKeyword(index string, field string, value interface{}, limit int) *[]SearchResult
	CreateNewIndex(index string) error
	CreateNewDocumentInIndex(index string, document *Document) (*Document ,error)
	GetDocumentById(index string, id string) (*Document, error)
}

// SearchResult used to gather results of elastic search calls
type SearchResult struct {
	ID   interface{}
	Body interface{}
}

type Document struct {
	ID uint64
	Body interface{}
}

type Article struct {
	Name string
}

// SearchWithKeyword allow you to search in elastic search by passing the desired index and keywords.
// Calls mapSearchResults.
// Returns a slice of SearchResult struct or nil.
func SearchWithKeyword(index string, field string, value interface{}, limit int) *[]SearchResult {
	client := database.EsClient
	terms := elastic.NewTermQuery(field, value)
	response, err := client.Search().
		Index(index).
		Query(terms).
		From(0).
		Size(limit).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil
	}
	prettifyNotFoundError(response)
	var article Article
	for _, item := range response.Each(reflect.TypeOf(article)) {
		if t, ok := item.(Article); ok {
			log.Info("Article by %s", t.Name)
		}
	}
	//return mapSearchResults(response)
	return nil
}

// prettifyNotFoundError returns an human readable error with a reason and type.
func prettifyNotFoundError(response *elastic.SearchResult)  {
	if response.Error != nil {
		log.Error("Error on search query: %v", response)
	}
}

// mapSearchResults map the results gathered from elastic search call and map them to a slice of SearchResult struct.
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

// CreateNewIndex allows you to create a new elastic search index.
// Checks if index already exists
func CreateNewIndex(index string) error {
	client := database.EsClient
	ctx := context.Background()
	exists, err := client.IndexExists(index).Do(ctx)
	if !exists {
		log.Error("Index exist already: %s", index)
		return errors.New("index exist already")
	}
	indexed, err := client.CreateIndex(index).Do(ctx)
	if err != nil {
		log.Error("cannot create new index: %s", index)
		return errors.New("cannot create new index: " + index)
	}
	if !indexed.Acknowledged{
		log.Error("error while acknowledging index: %s", index)
		return errors.New("error while acknowledging index: " + index)
	}
	return nil
}

func CreateNewDocumentInIndex(index string, document *Document) (*Document ,error) {
	client := database.EsClient
	inserted, err := client.Index().
		Index(index).
		BodyJson(document.Body).
		Do(context.Background())

	if err != nil {
		log.Error("cannot insert document in index %s", index)
		return nil, errors.New( "cannot insert document in index: " + index)
	}
	document.ID = helpers.ParseStringToUint64(inserted.Id)

	return document, nil
}

func GetDocumentById(index string, id string) (*Document, error) {
	client := database.EsClient
	ctx := context.Background()
	document, err := client.Get().
		Index(index).
		Id(id).
		Do(ctx)

	if err != nil {
		log.Error("cannot fetch document with id: %s", id)
		return nil, errors.New( "cannot fetch document with id: " + id)
	}
	if !document.Found {
		log.Error("document not found with id: %s", id)
		return nil, errors.New( "document not found with id: " + id)
	}
	_, err = client.Flush().Index(index).Do(ctx)
	if err != nil {
		log.Error("error while writing document")
		return nil, errors.New( "error while writing document")
	}
	parsedId := helpers.ParseStringToUint64(document.Id)
	doc := Document{
		ID: parsedId,
		Body: document.Source,
	}

	return &doc, nil
}