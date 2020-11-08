package router

import (
	v1 "server/api/v1"
	"github.com/gin-gonic/gin"
)
// 文章相关路由
func InitArticles(Router *gin.RouterGroup)(R gin.IRouter)  {
	ArticleRouter:=Router.Group("article")
	{
		ArticleRouter.POST("articleList",v1.ArticleList) //文章列表
		ArticleRouter.GET("articleDetail",v1.GetArticleDetail) // 文章详情
		ArticleRouter.GET("articlePassword",v1.CheckPassword)//验证文章密码是否正确
	    ArticleRouter.GET("articleClassify",v1.ArticleClassify) //文章分类
	    ArticleRouter.GET("hotArticle",v1.HotArticle)//热门文章
	    ArticleRouter.GET("tag",v1.AllTag)//tag获取
	    ArticleRouter.GET("config",v1.GetConfig)//获取博客配置
	    ArticleRouter.GET("github",v1.GetGithub)
	}
	return ArticleRouter
}