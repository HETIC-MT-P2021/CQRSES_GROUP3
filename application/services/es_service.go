package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP/core/es"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type EsConnector struct {
}

type EsService interface {
	SearchWithKeyword(index string, query map[string]interface{})
	CreateNewIndex(index string) error
	CreateNewDocumentInIndex(index string, document *Document) (*Document ,error)
	GetDocumentById(index string, id string) (*Document, error)
}

type SearchResult struct {
	ID   interface{}
	Body interface{}
}

type Document struct {
	ID   string
	Body es.Event
}

func SearchWithKeyword(index string, query *map[string]interface{}) *[]SearchResult {
	client, err := database.GetOriginalESClient()
	if err != nil {
		log.Error("cannot get elastic original client.")
		return nil
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Error("Error encoding query: %s", err)
	}
	response, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(index),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
		client.Search.WithSize(10),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	prettifyNotFoundError(response)

	return mapSearchResults(response)
}

func prettifyNotFoundError(response *esapi.Response)  {
	if response.IsError() {
		var error map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&error); err != nil {
			log.Fatal("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatal("[%s] %s: %s",
				response.Status(),
				error["error"].(map[string]interface{})["type"],
				error["error"].(map[string]interface{})["reason"],
			)
		}
	}
}

func mapSearchResults(response *esapi.Response) *[]SearchResult {
	var body map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		log.Fatal("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		response.Status(),
		int(body["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(body["took"].(float64)),
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
	if exists {
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

func CreateNewDocumentInIndex(index string, document *Document) error {
	client := database.EsClient
	inserted, err := client.Index().
		Index(index).
		BodyJson(document.Body).
		Do(context.Background())

	if err != nil {
		log.Error("cannot insert document in index %s", index)
		return errors.New( "cannot insert document in index: " + index)
	}
	log.Info(helpers.ParseStringToUint64(inserted.Id))
	document.ID = inserted.Id

	return nil
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
	doc := Document{
		ID: document.Id,
		Body: document.Source,
	}

	return &doc, nil
}

func constructQuery(keyword string, size int) *strings.Reader {
	var query = `{"query": {`
	query = query + keyword
	query = query + `}, "size": ` + strconv.Itoa(size) + `}`
	log.Info("\nquery:", query)
	isValid := json.Valid([]byte(query)) // returns bool
	if isValid == false {
		log.Info("query string not valid:", query)
		log.Info("Using default match_all query")
		query = "{}"
	} else {
		log.Info("valid JSON:", isValid)
	}
	var b strings.Builder
	b.WriteString(query)
	read := strings.NewReader(b.String())

	return read
}
