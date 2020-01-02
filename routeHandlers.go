package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sammarth-kapse/FileDownloadManager/repository"
	"log"
	"net/http"
)

func getHealthCheck(ctx *gin.Context) {

	ctx.String(http.StatusOK, "OK")
}

func downloadFiles(ctx *gin.Context) {

	var downloadRequest repository.DownloadRequest
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

	id := ctx.Param("downloadID")

	if response, ok := getDownloadInformationByID(id); ok {

		jsonFiles, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id":            response.ID,
			"start_time":    response.StartTime,
			"end_time":      response.EndTime,
			"status":        response.Status,
			"download_type": response.DownloadType,
			"files":         string(jsonFiles),
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"internal_code": 4002,
			"message":       "unknown download ID",
		})
	}

}
