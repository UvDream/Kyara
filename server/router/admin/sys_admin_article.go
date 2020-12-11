package admin

import (
	v1 "server/api/v1/admin"

	"github.com/gin-gonic/gin"
)

func InitAdminArticle(Router *gin.RouterGroup) (R gin.IRouter) {
	AdminArticleRouter := Router.Group("admin")
	{
		AdminArticleRouter.POST("add", v1.AddArticle)                // 添加文章
		AdminArticleRouter.GET("articleDetail", v1.GetArticleDetail) //查询文章详情
		AdminArticleRouter.POST("addNotice", v1.AddNotice)           //增加公告
		AdminArticleRouter.GET("addTag", v1.AddTag)                  //增加tag
		AdminArticleRouter.POST("blogComment", v1.GetBlogComment)    //获取博客留言
		AdminArticleRouter.POST("checkComment", v1.CheckBlogComment) //审核留言
		AdminArticleRouter.POST("editClassify",v1.EditClassify)//新增或者修改分类
		AdminArticleRouter.GET("deleteClassify",v1.DeleteClassify)//删除分类
	}
	return AdminArticleRouter
}
