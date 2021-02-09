package main

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/routes"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/helpers"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Connect to database and execute migrations
	cfg := database.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	err := database.Init(cfg)
	helpers.DieOnError("database connection failed", err)
	database.Migrate()

	// Setup router
	router := gin.Default()
	ecfg := database.EsCfg{Url: "http://es:9200"}
	database.GetESClient(&ecfg)
	if err := services.MigrateIndex(); err != nil {
		helpers.DieOnError("migration failed", err)
	}

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "Authorization",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	routes.Init(router)
	domain.InitBuses()

	go func() {
		if err := router.Run(":8000"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// ----------------- CLOSE APP -----------------

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")
}
