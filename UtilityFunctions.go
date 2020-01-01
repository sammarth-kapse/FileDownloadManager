package main

import (
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func getDownloadResponse(downloadRequest DownloadRequest) string {

	if isNonValidType(downloadRequest.Type) {
		return ""
	}

	currID := uuid.New().String()
	currStatus := new(DownloadStatus)
	downloadCollection[currID] = currStatus

	createDirectory(FILE_PATH + currID)

	if downloadRequest.Type == "serial" {
		serialDownloader(downloadRequest.Urls, currID)
	} else {
		go concurrentDownloader(downloadRequest.Urls, currID)
	}

	return currID
}

func isNonValidType(downloadType string) bool {
	if strings.Compare(downloadType, "serial") == 0 {
		return false
	}
	if strings.Compare(downloadType, "concurrent") == 0 {
		return false
	}
	return true
}

func (downloadStatus *DownloadStatus) initializeStatus(id, downloadType, status string) {

	downloadStatus.Id = id
	downloadStatus.Status = status
	downloadStatus.DownloadType = downloadType
	downloadStatus.Files = make(map[string]string)
	downloadStatus.StartTime = time.Now()
}

func serialDownloader(urls []string, currID string) {

	downloadCollection[currID].initializeStatus(currID, "SERIAL", "QUEUED")

	for _, v := range urls {

		destination := FILE_PATH + currID + "/" + getFileName(v)
		err := downloadFile(v, destination)
		if err != nil {
			downloadCollection[currID].Status = "FAILED"
			downloadCollection[currID].EndTime = time.Now()
			log.Fatal(err)
		}
		downloadCollection[currID].Files[v] = destination
	}

	downloadCollection[currID].EndTime = time.Now()
	downloadCollection[currID].Status = "SUCCESSFUL"
}

func concurrentDownloader(urls []string, currID string) {

	downloadCollection[currID].initializeStatus(currID, "CONCURRENT", "QUEUED")
	var wg sync.WaitGroup

	for _, v := range urls {

		destination := FILE_PATH + currID + "/" + getFileName(v)
		wg.Add(1)
		go concurrentDownloadHelper(v, destination, currID, &wg)
		downloadCollection[currID].Files[v] = destination
	}
	wg.Wait()

	downloadCollection[currID].EndTime = time.Now()
	downloadCollection[currID].Status = "SUCCESSFUL"
}

func concurrentDownloadHelper(url, filepath, currID string, wg *sync.WaitGroup) {
	err := downloadFile(url, filepath)
	if err != nil {
		downloadCollection[currID].Status = "FAILED"
		downloadCollection[currID].EndTime = time.Now()
		log.Fatal(err)
	}
	wg.Done()
}

func createDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getFileName(url string) string {
	return uuid.New().String()
}

func downloadFile(url string, filepath string) error {

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
