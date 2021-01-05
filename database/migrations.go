package database

import (
	"github.com/edwinvautier/go-boilerplate/models"
	log "github.com/sirupsen/logrus"
)

// Migrate executes migrations once the db is connected
func Migrate() {
	log.Info("Executing migrations...")
	Db.AutoMigrate(&models.Customer{})
}