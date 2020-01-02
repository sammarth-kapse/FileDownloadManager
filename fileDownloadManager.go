package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	initializeRoutes(router)

	router.Run(":8081")
}
