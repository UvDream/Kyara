package initialize

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
	"server/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RoutersGroupApp.System
	//跨域设置
	Router.Use(middleware.Cors()) //放行所有的请求
	//Router.Use(middleware.CorsByRules())  //按照配置规则放行跨域
	//公开的接口
	PublicGroup := Router.Group("/public")
	{
		PublicGroup.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "程序正常运行",
			})
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}
	//需要登陆才可以访问的接口
	PrivateGroup := Router.Group("")
	{
		systemRouter.InitUserRouter(PrivateGroup)
	}
	return Router
}
