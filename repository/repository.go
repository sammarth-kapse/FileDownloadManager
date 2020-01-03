package repository

import "time"

var downloadCollection = make(map[string]*DownloadInformation)

func InsertIntoDownloadCollection(id string, information *DownloadInformation) {

	downloadCollection[id] = information
}

func GetDownloadInformationByID(id string) (*DownloadInformation, bool) {

	downloadItem, ok := downloadCollection[id]

	if ok {
		return downloadItem, ok
	}
	return nil, ok
}

func setStartTimeForGivenID(id string, startTime time.Time) {

	downloadCollection[id].StartTime = startTime
}

func setEndTimeForGivenID(id string, endTime time.Time) {

	downloadCollection[id].EndTime = endTime
}

func setStatusForGivenID(id string, status string) {

	downloadCollection[id].Status = status
}

func appendFileForGivenID(id, url, filePath string) {

	downloadCollection[id].Files[url] = filePath
}
