package admin

import (
	v1 "server/api/v1/admin"
	"server/api/v1/blog"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func InitAdminArticle(Router *gin.RouterGroup) (R gin.IRoutes) {
	AdminArticleRouter := Router.Group("admin").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		AdminArticleRouter.POST("add", v1.AddArticle)                      // 添加文章
		AdminArticleRouter.GET("articleDetail", v1.GetArticleDetail)       //查询文章详情
		AdminArticleRouter.POST("addNotice", v1.AddNotice)                 //增加公告
		AdminArticleRouter.GET("deleteNotice",v1.DeleteNotice) 				//删除公告
		AdminArticleRouter.GET("addTag", v1.AddTag)                        //增加tag
		AdminArticleRouter.POST("blogComment", v1.GetBlogComment)          //获取博客留言
		AdminArticleRouter.POST("checkComment", v1.CheckBlogComment)       //审核留言
		AdminArticleRouter.POST("editClassify", v1.EditClassify)           //新增或者修改分类
		AdminArticleRouter.GET("deleteClassify", v1.DeleteClassify)        //删除分类
		AdminArticleRouter.GET("iconfontClassify", v1.GetIconfontClassify) //获取阿里矢量分类图标
		AdminArticleRouter.POST("revertComment", v1.RevertComment)         //回复留言
		AdminArticleRouter.GET("deleteComment", v1.DeleteComment)          //删除留言
		AdminArticleRouter.POST("saveConfig",v1.SaveConfig)					//保存设置
		AdminArticleRouter.POST("addInterviewClassify",blog.AddInterviewClassify)		//新增面试题分类
		AdminArticleRouter.GET("deleteInterviewClassify",blog.DeleteInterviewClassify)//删除面试题分类
		AdminArticleRouter.POST("addInterview",blog.AddInterview)//新增面试题
	}
	return AdminArticleRouter
}
