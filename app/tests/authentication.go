package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	log "github.com/sirupsen/logrus"
)

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
		cookieToken = strings.Split(cookieToken, ";")[0]

		return statusOK && pageOK
	})

	return cookieToken
}
