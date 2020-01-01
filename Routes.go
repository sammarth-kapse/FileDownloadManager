package main

func initializeRoutes() {

	router.GET("/health", getHealthCheck)
	downloadRoutes := router.Group("/downloads")
	{
		downloadRoutes.POST("/", Downloader)

		downloadRoutes.GET("/:downloadID", getDownloadStatus)
	}
	router.GET("/files", getDownloadedFiles)
}
