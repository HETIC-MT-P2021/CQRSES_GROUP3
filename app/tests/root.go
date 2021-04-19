package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func testRoot(t *testing.T, token string) {
	req, _ := createRequest("GET", "/api/", nil)
	log.Info("token used ", token)
	req.Header.Set("Authorization", "Bearer "+token)
	testHTTPResponse(t, router, req, "error at api root", func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			log.Error("read body error :", err)
		}
		pageOK := err == nil && strings.Contains(string(p), "OKKKKKK")
		if !pageOK {
			log.Error(string(p))
		}
		return statusOK && pageOK
	})
}
