package initialize

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/middleware"
	"server/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//Router.Static("/files", "./files")
	//设置静态文件夹
	Router.StaticFS("/files", gin.Dir(global.Config.Local.Path, true))
	//系统相关路由
	systemRouter := router.RoutersGroupApp.System
	//文章相关路由
	articleRouter := router.RoutersGroupApp.Article
	//标签相关路由
	tagRouter := router.RoutersGroupApp.Tag
	//分类相关路由
	category := router.RoutersGroupApp.Category
	//文件相关路由
	file := router.RoutersGroupApp.File
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
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitSysRouter(PrivateGroup)
		articleRouter.InitArticleRouter(PrivateGroup)
		tagRouter.InitTagRouter(PrivateGroup)
		category.InitCategoriesRouter(PrivateGroup)
		file.InitFileRouter(PrivateGroup)
	}
	return Router
}
