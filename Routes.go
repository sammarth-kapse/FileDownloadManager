package main

import . "github.com/sammarth-kapse/FileDownloadManager/DownloadHandler"

func initializeRoutes() {

	Router.GET("/health", getHealthCheck)
	downloadRoutes := Router.Group("/downloads")
	{
		downloadRoutes.POST("/", downloader)

		downloadRoutes.GET("/:downloadID", getDownloadStatus)
	}
	Router.GET("/files", getDownloadedFiles)
}
