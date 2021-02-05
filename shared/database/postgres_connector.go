package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	// postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Db is the database object
var Db *gorm.DB

// Config is the structure used to load db credentials from the environment.
type Config struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
}

// Init Initializes a db connection
func Init(cfg Config) error {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	var tmpDb *gorm.DB
	var err error

	// Try connecting database 5 times
	for test := 1; test <= 5; test++ {
		tmpDb, err = gorm.Open("postgres", dbURL)

		if err != nil {
			log.Warn("db connection failed. (%s/5)", test)
			time.Sleep(5 * time.Second)
			continue
		}

		break
	}
	if err != nil {
		return err
	}

	Db = tmpDb
	log.Info("Connected to database!")

	return nil
}
