package main

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {

	router.GET("/health", getHealthCheck)

	downloadRoutes := router.Group("/downloads")
	{
		downloadRoutes.POST("/", processDownloading)

		downloadRoutes.GET("/:downloadID", getDownloadStatus)
	}
}
