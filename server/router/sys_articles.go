package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)
// 文章相关路由
func InitArticles(Router *gin.RouterGroup)(R gin.IRouter)  {
	ArticleRouter:=Router.Group("article")
	{
		ArticleRouter.POST("articleList",v1.ArticleList) //文章列表
		ArticleRouter.GET("articleDetail",v1.GetArticleDetail) // 文章详情
		ArticleRouter.GET("articlePassword",v1.CheckPassword)//验证文章密码是否正确
	}
	return ArticleRouter
}