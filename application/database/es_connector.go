package database

import (
	elastic "github.com/elastic/go-elasticsearch/v8"
	log "github.com/sirupsen/logrus"
)

// GetESClient Attempts to connect to elastic search
// Return the client and error
func GetESClient() (*elastic.Client, error) {
	cfg := elastic.Config{
		Addresses: []string{
			"http://es:9200",
		},
	}

	es, err := elastic.NewClient(cfg)
	if err != nil {
		log.Error("Error creating the client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Error("Error gettingg response: %s", err)
	}

	if err = res.Body.Close(); err != nil {
		log.Error("Error closing connection: %s", err)
	}

	return es, err
}
