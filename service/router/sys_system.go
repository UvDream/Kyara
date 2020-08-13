package router

import (
	"blog-api/api/v1"
	"blog-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("system").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.POST("getSystemConfig", v1.GetSystemConfig) // 获取配置文件内容
		UserRouter.POST("setSystemConfig", v1.SetSystemConfig) // 设置配置文件内容
	}
}
