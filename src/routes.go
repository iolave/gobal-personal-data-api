package main

import (
	"github.com/iolave/gobal-personal-data-api/humans"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.GET("/", humans.test)

}

func test(c *gin.Context) {
	a := humans.HumanId{
		Country:      "CL",
		DocumentType: "RUT",
		DocumentId:   "19408138-3",
	}

	c.JSON(http.StatusOK, a)
}
