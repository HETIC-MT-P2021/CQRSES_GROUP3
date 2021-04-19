package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/routes"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	"github.com/gin-gonic/gin"
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

	return http.NewRequest(method, "http://localhost:8000" + route, bytesBody)
}

func TestAuthentication(t *testing.T) {
	initDB()
	initRouter()

	credentials := models.CustomerForm {
		Email: "john@doe.com",
		Name: "John Doe",
		Password: "Test123456",
	}

	testRegister(t, credentials)
	token := testLogin(t, credentials)
	log.Info(token)
}

func testRegister(t *testing.T, credentials models.CustomerForm) {
	req, _ := createRequest("POST", "/register", credentials)
	
	testHTTPResponse(t, router, req, "couldn't register customer", func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			log.Error("read body error :", err)
		}
		pageOK := err == nil && strings.Index(string(p), "John Doe") > 0

		return statusOK && pageOK
	})
}

func testLogin(t *testing.T, credentials models.CustomerForm) string {
	req, _ := createRequest("POST", "/login", credentials)
	var cookieToken string
	testHTTPResponse(t, router, req, "Login failed", func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			log.Error("read body error :", err)
		}
		pageOK := err == nil && strings.Contains(string(p), "Logged in successfully")
		if !pageOK {
			log.Error("received : ", string(p))
		}
		cookieToken = strings.Split(w.Header().Get("Set-Cookie"), "=")[1]

		return statusOK && pageOK
	})

	return cookieToken
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