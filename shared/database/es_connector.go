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

// GetESClient Attempts to connect to elastic search using "github.com/olivere/elastic/v7" package.
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

// GetOriginalESClient Attempts to connect to elastisearch service using "github.com/elastic/go-elasticsearch/v8" package.
// Return the client and error
func GetOriginalESClient() (*es.Client, error) {
	cfg := es.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	elsrch, err := es.NewClient(cfg)
	if err != nil {
		log.Errorf("Error creating the client: %s", err)
	}
	res, err := elsrch.Info()
	if err != nil {
		log.Errorf("Error getting response: %s", err)
	}

	if err = res.Body.Close(); err != nil {
		log.Errorf("Error closing connection: %s", err)
	}
	log.Info("Connected to elasticsearch")

	return elsrch, err
}
