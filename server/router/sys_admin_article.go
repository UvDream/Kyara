package router

import (
	v1 "gin-vue-admin/api/v1/admin"
	"github.com/gin-gonic/gin"
)

func InitAdminArticle(Router *gin.RouterGroup)(R gin.IRouter)  {
	AdminArticleRouter:=Router.Group("adminArticle")
	{
		AdminArticleRouter.POST("add",v1.AddArticle) // 添加文章
	}
	return AdminArticleRouter
}