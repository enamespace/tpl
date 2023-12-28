package user

import (
	"log"

	"github.com/gin-gonic/gin"

	v1 "github.com/enamespace/tpl/internal/api/v1"
	"github.com/enamespace/tpl/pkg/core"
	"github.com/enamespace/tpl/pkg/errors"
)

func (u *UserController) Create(c *gin.Context) {
	log.Println("create user function called.")
	var r v1.User

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	if errs := r.Validate(); len(errs) > 0 {
		core.WriteResponse(c, errors.NewAggregate(errs), nil)
		return
	}

	err := u.srv.User().Create(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, r)
}
