package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {

		c.JSON(http.StatusInternalServerError, nil)

		return
	}

	c.JSON(http.StatusOK, data)
}
