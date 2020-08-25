package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)
// 文章相关路由
func InitArticles(Router *gin.RouterGroup)(R gin.IRouter)  {
	ArticleRouter:=Router.Group("article")
	{
		ArticleRouter.POST("articleList",v1.ArticleList)
	}
	return ArticleRouter
}