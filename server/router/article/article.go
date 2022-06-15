package article

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type ArticlesRouter struct {
}

func (article *ArticlesRouter) InitArticleRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	articleRouter := Router.Group("article")
	articleApi := v1.ApiGroupApp.ArticleApiGroup
	{
		articleRouter.POST("/create", articleApi.CreateArticle)
		articleRouter.DELETE("/delete", articleApi.DeleteArticle)
		articleRouter.PUT("/update", articleApi.UpdateArticle)
		articleRouter.POST("/list", articleApi.GetArticleList)
		articleRouter.GET("/history", articleApi.GetArticleHistory)
	}
	return articleRouter
}
