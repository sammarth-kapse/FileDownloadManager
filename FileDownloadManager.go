package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/sammarth-kapse/FileDownloadManager/DownloadHandler"
)

func main() {

	Router = gin.Default()
	initializeRoutes()
	Router.Run(":8081")
}
