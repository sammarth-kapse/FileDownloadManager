package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sammarth-kapse/FileDownloadManager/repository"
	"log"
	"net/http"
)

func getHealthCheck(ctx *gin.Context) {

	ctx.String(http.StatusOK, "OK")
}

func processDownloading(ctx *gin.Context) {

	var downloadRequest repository.DownloadRequest

	// Request's Body(in JSON format) stored in downloadRequest
	err := ctx.BindJSON(&downloadRequest)
	if err != nil {
		log.Fatal(err)
		return
	}

	if !isValidType(downloadRequest.Type) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"internal_code": 4001,
			"message":       "unknown type of download",
		})
	} else if isURLsEmpty(downloadRequest.URLs) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"internal_code": 4003,
			"message":       "no files to download",
		})
	} else {
		id := getDownloadID(downloadRequest)
		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

func getDownloadStatus(ctx *gin.Context) {

	// Get id from parameters in given path
	id := ctx.Param("downloadID")

	if response, isPresent := getDownloadInformationByID(id); isPresent {

		ctx.JSON(http.StatusOK, gin.H{
			"id":            response.ID,
			"start_time":    response.StartTime,
			"end_time":      response.EndTime,
			"status":        response.Status,
			"download_type": response.DownloadType,
			"files":         response.Files,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"internal_code": 4002,
			"message":       "unknown download ID",
		})
	}
}
