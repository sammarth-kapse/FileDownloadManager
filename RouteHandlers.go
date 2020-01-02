package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/sammarth-kapse/FileDownloadManager/DownloadHandler"
	"log"
	"net/http"
)

func getHealthCheck(ctx *gin.Context) {

	ctx.String(http.StatusOK, "OK")
}

func downloadFiles(ctx *gin.Context) {

	var downloadRequest DownloadRequest
	err := ctx.BindJSON(&downloadRequest)
	if err != nil {
		log.Fatal(err)
		return
	}

	id := GetDownloadResponse(downloadRequest)
	// id = "" (empty string) is received for a bad request
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"internal_code": 4001,
			"message":       "unknown type of download",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

func getDownloadStatus(ctx *gin.Context) {

	id := string(ctx.Param("downloadID"))

	if response, ok := DownloadCollection[id]; ok {
		jsonFiles, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"id":            response.Id,
			"start_time":    response.StartTime,
			"end_time":      response.EndTime,
			"status":        response.Status,
			"download_type": response.DownloadType,
			"files":         string(jsonFiles),
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"internal_code": 4002,
			"message":       "unknown DownloadHandler ID",
		})
	}

}

func getDownloadedFiles(ctx *gin.Context) {

}
