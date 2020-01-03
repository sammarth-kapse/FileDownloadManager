package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var testURLs = []string{
	"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg",
	"https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG",
}

func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) *httptest.ResponseRecorder {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
	return w
}

// To check if some pattern is present in given body
func testIsPresentInResponse(body, matchingPattern string) bool {

	if strings.Compare(body, matchingPattern) >= 0 {
		return true
	}
	return false
}
