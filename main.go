package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	initializeRoutes()
	// Start serving the application
	router.Run()
}
 