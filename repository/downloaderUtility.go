package repository

import (
	"log"
	"os"
	"time"
)

type DownloadRequest struct {
	Type string   `json:type`
	URLs []string `json:urls`
}

var GLOBAL_PATH string = "/Users/sammarthkapse/Downloads/goDownloads/"
var QUEUED = "QUEUED"
var SERIAL = "SERIAL"
var CONCURRENT = "CONCURRENT"
var FAILED = "FAILED"
var SUCCESS = "SUCCESSFUL"

func (information DownloadInformation) createDirectory() {

	createDirectoryIfNotExist(information.DirectoryPath)
}

func createDirectoryIfNotExist(directoryPath string) {

	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		err = os.MkdirAll(directoryPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Marks the start of particular download request
func (information DownloadInformation) markDownloadStart() {

	information.StartTime = time.Now()
	information.Status = QUEUED

	setStartTimeForGivenID(information.ID, information.StartTime)
	setStatusForGivenID(information.ID, information.Status)
}

// Marks the end of particular download request
func (information DownloadInformation) markDownloadEnd(status string) {

	information.EndTime = time.Now()
	information.Status = status

	setEndTimeForGivenID(information.ID, information.EndTime)
	setStatusForGivenID(information.ID, information.Status)
}

func (information DownloadInformation) appendDownloadFile(url, filePath string) {

	information.Files[url] = filePath

	appendFileForGivenID(information.ID, url, filePath)
}

func getFileNameFromURL(url string) string {

	var fileName []byte
	lastSlashPosition := -1
	for pos, val := range url {
		if val == '/' {
			lastSlashPosition = pos
		}
	}

	// Appends all bytes(characters) of url after last '/' into fileName
	for pos := lastSlashPosition + 1; pos != len(url); pos++ {
		fileName = append(fileName, url[pos])
	}
	return string(fileName)
}
