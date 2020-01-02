package main

import (
	"github.com/sammarth-kapse/FileDownloadManager/repository"
	"strings"
)

var SERIAL = "SERIAL"
var CONCURRENT = "CONCURRENT"

func getDownloadID(downloadRequest repository.DownloadRequest) string {

	downloadInformation := repository.New(downloadRequest)

	downloadInformation.Download()

	return downloadInformation.ID
}

func getDownloadInformationByID(id string) (*repository.DownloadInformation, bool) {

	return repository.GetDownloadInformationByID(id)
}

func isValidType(downloadType string) bool {

	if strings.Compare(downloadType, SERIAL) == 0 {
		return true
	}
	if strings.Compare(downloadType, CONCURRENT) == 0 {
		return true
	}
	return false
}

func isURLsEmpty(URLs []string) bool {

	return len(URLs) == 0
}
