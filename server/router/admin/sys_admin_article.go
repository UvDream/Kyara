package admin

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1/admin"
)

func InitAdminArticle(Router *gin.RouterGroup)(R gin.IRouter)  {
	AdminArticleRouter:=Router.Group("adminArticle")
	{
		AdminArticleRouter.POST("add",v1.AddArticle) // 添加文章
		AdminArticleRouter.GET("articleDetail",v1.GetArticleDetail)//查询文章详情
	}
	return AdminArticleRouter
}