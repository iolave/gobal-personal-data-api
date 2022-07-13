package humans

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func myFunc() int {
	fmt.Println("test")
	return (1)
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	// })
}

func test(c *gin.Context) {
	a := HumanId{
		Country:      "CL",
		DocumentType: "RUT",
		DocumentId:   "19408138-3",
	}

	c.JSON(http.StatusOK, a)
}
