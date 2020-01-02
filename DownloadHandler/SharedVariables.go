package DownloadHandler

import (
	"github.com/gin-gonic/gin"
	"time"
)

type DownloadRequest struct {
	Type string   `json:type`
	Urls []string `json:urls`
}

type DownloadStatus struct {
	Id           string            `json:"id"`
	StartTime    time.Time         `json:"startTime"`
	EndTime      time.Time         `json:"endTime"`
	Status       string            `json:"status"`
	DownloadType string            `json:"downloadType"`
	Files        map[string]string `json:"files"`
}

// Collection of all the download-status mapped through download-id
var DownloadCollection = make(map[string]*DownloadStatus)

// Location where all the downloaded files are located
var GLOBAL_PATH string = "/Users/sammarthkapse/Downloads/goDownloads/"

var Router *gin.Engine
