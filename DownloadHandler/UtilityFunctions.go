package DownloadHandler

import (
	"log"
	"os"
	"strings"
	"time"
)

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

func createDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getFileName(url string) string {
	var fileName []byte
	lastSlash := -1
	for i, v := range url {
		if v == '/' {
			lastSlash = i
		}
	}
	for i := lastSlash + 1; i != len(url); i++ {
		fileName = append(fileName, url[i])
	}
	return string(fileName)
}
