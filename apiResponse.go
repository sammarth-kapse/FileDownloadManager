package main

import "github.com/sammarth-kapse/FileDownloadManager/repository"

func getDownloadID(downloadRequest repository.DownloadRequest) string {

	downloadInformation := repository.New(downloadRequest)

	downloadInformation.Download()

	return downloadInformation.ID
}

func getDownloadInformationByID(id string) (*repository.DownloadInformation, bool) {

	return repository.GetDownloadInformationByID(id)
}
