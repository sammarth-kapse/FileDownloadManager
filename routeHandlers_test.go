package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sammarth-kapse/FileDownloadManager/repository"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHealthCheck(t *testing.T) {

	router := gin.Default()
	router.GET("/health", getHealthCheck)
	request, _ := http.NewRequest("GET", "/health", nil)

	testHTTPResponse(t, router, request, func(w *httptest.ResponseRecorder) bool {

		statusCodeOK := w.Code == http.StatusOK

		response, err := ioutil.ReadAll(w.Body)

		return statusCodeOK && err == nil && testIsPresentInResponse(string(response), "OK")
	})
}

func TestProcessDownloading(t *testing.T) {

	tests := []struct {
		downloadRequest repository.DownloadRequest
		expectedCode    int
	}{
		{repository.DownloadRequest{SERIAL, testURLs}, http.StatusOK},
		{repository.DownloadRequest{CONCURRENT, testURLs}, http.StatusOK},
		{repository.DownloadRequest{"SERIALXY", testURLs}, http.StatusBadRequest},   // Non-Valid Download-Type
		{repository.DownloadRequest{"", testURLs}, http.StatusBadRequest},           // Non-valid Download-Type
		{repository.DownloadRequest{SERIAL, []string{}}, http.StatusBadRequest},     // Empty URL List
		{repository.DownloadRequest{CONCURRENT, []string{}}, http.StatusBadRequest}, // Empty URL List
	}

	for _, test := range tests {

		router := gin.Default()
		router.POST("/downloads", processDownloading)

		jsonRequest, err := json.Marshal(test.downloadRequest)
		if err != nil {
			log.Fatal(err)
		}

		request, _ := http.NewRequest("POST", "/downloads", bytes.NewBuffer(jsonRequest))

		testHTTPResponse(t, router, request, func(w *httptest.ResponseRecorder) bool {
			return test.expectedCode == w.Code
		})
	}
}

func TestGetDownloadStatus(t *testing.T) {

	nonInsertedID := uuid.New().String()
	insertedSerialID := testGetIDForDownloadRequest(repository.DownloadRequest{SERIAL, testURLs})
	insertedConcurrentID := testGetIDForDownloadRequest(repository.DownloadRequest{CONCURRENT, testURLs})

	tests := []struct {
		id           string
		expectedCode int
	}{
		{insertedConcurrentID, http.StatusOK},
		{insertedSerialID, http.StatusOK},
		{"rfgerug", http.StatusBadRequest},     // ID does not exist in downloadCollection
		{nonInsertedID, http.StatusBadRequest}, // ID not inserted in downloadCollection
		{"", http.StatusNotFound},              // Wrong Request Format
	}

	for _, test := range tests {

		router := gin.Default()
		router.GET("/downloads/:downloadID", getDownloadStatus)

		request, _ := http.NewRequest("GET", "/downloads/"+test.id, nil)

		testHTTPResponse(t, router, request, func(w *httptest.ResponseRecorder) bool {
			return test.expectedCode == w.Code
		})
	}
}

// To process a download-request and return the obtained ID for corresponding request
func testGetIDForDownloadRequest(downloadRequest repository.DownloadRequest) string {

	router := gin.Default()

	jsonBody, err := json.Marshal(downloadRequest)
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/downloads", processDownloading)
	request, _ := http.NewRequest("POST", "/downloads", bytes.NewBuffer(jsonBody))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	jsonResponse := w.Body.Bytes()

	var response struct {
		ID string
	}
	err = json.Unmarshal(jsonResponse, &response)
	if err != nil {
		log.Fatal(err)
	}

	return response.ID
}
