package internal

import (
	"github.com/enamespace/tpl/internal/controller/user"
	"github.com/enamespace/tpl/internal/pkg/middleware"
	"github.com/enamespace/tpl/internal/store/mysql"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	middleware.RequestID()

	v1 := g.Group("v1")

	storeIns, _ := mysql.GetMysqlFactory()

	{
		userv1 := v1.Group("user")

		{
			userController := user.NewUserController(storeIns)
			userv1.GET(":name", userController.Get)
		}
	}

}
