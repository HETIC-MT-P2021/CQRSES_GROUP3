package database

import (
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"time"
)

type EsCfg struct {
	Url string
}

var EsClient *elastic.Client

// GetESClient Attempts to connect to elastic search
// Return the client and error
func GetESClient(escfg *EsCfg) {
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL(escfg.Url),
		elastic.SetHealthcheckInterval(5*time.Second),
	)
	if err != nil {
		log.Error("Cannot connect to ")
	}

	EsClient = client
}
