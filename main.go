package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/caarlos0/env"
	"github.com/edwinvautier/go-boilerplate/database"
	"github.com/edwinvautier/go-boilerplate/helpers"
	"github.com/edwinvautier/go-boilerplate/routes"
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
