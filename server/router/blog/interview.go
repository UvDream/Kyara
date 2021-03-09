package blog

import (
	"github.com/gin-gonic/gin"
	"server/api/v1/blog"
)

func InitInterview(Router *gin.RouterGroup)(R gin.IRouter)  {
	InterviewRouter:=Router.Group("interview")
	{
		InterviewRouter.GET("classify",blog.GetInterviewClassify)//面试题分类接口
		InterviewRouter.POST("list",blog.GetInterview)//面试题列表接口
		InterviewRouter.GET("detail",blog.GetInterviewDetail)//面试题详情接口
	}
	return InterviewRouter

}