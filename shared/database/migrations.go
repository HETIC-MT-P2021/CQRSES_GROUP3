package database

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/repositories"
	log "github.com/sirupsen/logrus"
)

// Migrate executes migrations once the db is connected
func (r *repositories.Repository)Migrate() {
	log.Info("Executing migrations...")
	r.Db.AutoMigrate(&models.Customer{})
}
