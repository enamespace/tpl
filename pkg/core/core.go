package core

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.JSON(http.StatusOK, data)
}
