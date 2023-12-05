package user

import (
	srvv1 "github.com/enamespace/tpl/internal/service/v1"
	"github.com/enamespace/tpl/internal/store"
)

type UserController struct {
	srv srvv1.Service
}

func NewUserController(ds store.Factory) *UserController {
	return &UserController{
		srv: srvv1.NewService(ds),
	}
}
