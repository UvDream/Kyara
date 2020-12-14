package blog

import (
	"github.com/gin-gonic/gin"
	"server/api/v1/blog"
)

func InitOther(Router *gin.RouterGroup)(R gin.IRouter){
	OtherRouter:=Router.Group("other")
	{
		OtherRouter.POST("baidu",blog.ToBaiDu)//百度收录
		OtherRouter.GET("classifyStatistics",blog.GetClassifyStatistics)//获取分类动态
		OtherRouter.GET("blogDynamic",blog.GetBlogDynamic)//获取博客动态
	}
	return OtherRouter
}
