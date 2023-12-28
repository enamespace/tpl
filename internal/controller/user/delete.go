package user

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/enamespace/tpl/pkg/core"
)

func (u *UserController) Delete(c *gin.Context) {
	log.Println("delete user function called.")

	err := u.srv.User().Delete(c, c.Param("name"))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
