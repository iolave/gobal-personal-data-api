package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// test2
	router = gin.Default()
	initializeRoutes()
	router.Run()
}
