package humans

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type test interface {
	Test() gin.HandlerFunc
}

func (c *gin.Context) Test() gin.HandlerFunc {

	me := Human{
		Id: HumanId{
			Country:      "CHILE",
			DocumentType: "RUT",
			DocumentId:   "19408138",
		},
		Name:    "OLAVE RIQUELME IGNACIO ALONSO",
		Gender:  "MASCULINE",
		City:    "SANTIAGO",
		Address: "MI CASA 1",
	}
	fmt.Println(me.Address)

	c.JSON(200, me)

	return gin.HandlerFunc(fn)

}
