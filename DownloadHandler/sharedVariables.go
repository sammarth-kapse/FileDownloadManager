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

var DownloadCollection = make(map[string]*DownloadStatus)

var GLOBAL_PATH string = "/Users/sammarthkapse/Downloads/goDownloads/"

var Router *gin.Engine
