package user

import (
	"log"

	"github.com/enamespace/tpl/pkg/core"
	"github.com/gin-gonic/gin"
)

func (u *UserController) Get(c *gin.Context) {
	log.Println("get user function called.")

	user, err := u.srv.User().Get(c, c.Param("name"))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, user)
}
