package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/routes"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/helpers"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var router *gin.Engine

func initRouter() {
	router = gin.Default()
	routes.Init(router)
}

func initDB() {
	// Connect to database and execute migrations
	cfg := database.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	cfg.Name = "test-db"
	err := database.Init(cfg)
	helpers.DieOnError("database connection failed", err)
	database.Migrate()
	database.Db.Exec("DELETE FROM customers")
}

func createRequest(method string, route string, body interface{}) (*http.Request, error) {
	marshalledContent, err := json.Marshal(body)
	if err != nil {
		return &http.Request{}, err
	}
	bytesBody := bytes.NewBuffer(marshalledContent)

	return http.NewRequest(method, "http://localhost:8000"+route, bytesBody)
}
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, errorMessage string, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Errorf(errorMessage)
	}
}
