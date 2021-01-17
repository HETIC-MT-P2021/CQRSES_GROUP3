package database

import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"time"
)

type EsCfg struct {
	Url string
}

var EsClient *elastic.Client

// GetESClient Attempts to connect to elastic search from "github.com/olivere/elastic/v7"
// Affect the EsClient to elastic client.
func GetESClient(escfg *EsCfg) {
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL(escfg.Url),
		elastic.SetHealthcheckInterval(5*time.Second),
	)
	if err != nil {
		log.Error("Cannot connect to elastic search")
	} else {
		log.Info("Connected to elastic search")
	}

	EsClient = client
}

// GetOriginalESClient Attempts to connect to elastisearch service from "github.com/elastic/go-elasticsearch/v8"
// Return the client and error
func GetOriginalESClient() (*es.Client, error) {
	cfg := es.Config{
		Addresses: []string{
			"http://es:9200",
		},
	}

	elsrch, err := es.NewClient(cfg)
	if err != nil {
		log.Error("Error creating the client: %s", err)
	}
	res, err := elsrch.Info()
	if err != nil {
		log.Error("Error getting response: %s", err)
	}

	if err = res.Body.Close(); err != nil {
		log.Error("Error closing connection: %s", err)
	}
	log.Info("Connected to elasticsearch")

	return elsrch, err
}