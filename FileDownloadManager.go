package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router = gin.Default()
	initializeRoutes()
	router.Run(":8081")
}
