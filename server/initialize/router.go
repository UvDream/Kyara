package initialize

import (
	"github.com/gin-gonic/gin"
	"server/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RoutersGroupApp.System
	PrivateGroup := Router.Group("")
	{
		systemRouter.InitUserRouter(PrivateGroup)
	}
	return Router
}
