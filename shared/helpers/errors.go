package helpers

import (
	log "github.com/sirupsen/logrus"
)

// DieOnError checks if an error is nil and log fatal error message otherwise
func DieOnError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
