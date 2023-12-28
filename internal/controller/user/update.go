package user

import (
	"log"

	"github.com/gin-gonic/gin"

	v1 "github.com/enamespace/tpl/internal/api/v1"
	"github.com/enamespace/tpl/pkg/core"
	"github.com/enamespace/tpl/pkg/errors"
)

func (u *UserController) Update(c *gin.Context) {
	log.Println("udpate user function called.")
	var r v1.User

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	user, err := u.srv.User().Get(c, c.Param("name"))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	user.Email = r.Email
	user.NickName = r.NickName

	//TODO: validate for update
	if errs := user.Validate(); len(errs) > 0 {
		core.WriteResponse(c, errors.NewAggregate(errs), nil)
		return
	}

	err = u.srv.User().Update(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, r)
}
