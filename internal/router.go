package internal

import (
	"github.com/gin-gonic/gin"

	"github.com/enamespace/tpl/internal/controller/user"
	"github.com/enamespace/tpl/internal/pkg/middleware"
	"github.com/enamespace/tpl/internal/store/mysql"
)

func initRouter(g *gin.Engine) {
	g.Use(middleware.RequestID())

	v1 := g.Group("v1")

	storeIns, _ := mysql.GetMysqlFactory()

	{
		userv1 := v1.Group("user")

		{
			userController := user.NewUserController(storeIns)
			userv1.GET(":name", userController.Get)
			userv1.POST("", userController.Create)
			userv1.DELETE(":name", userController.Delete)
			userv1.PUT(":name", userController.Update)
		}
	}

}
