package blog

import "github.com/gin-gonic/gin"

func InitInterview(Router *gin.RouterGroup)(R gin.IRouter)  {
	InterviewRouter:=Router.Group("interview")
	{
		InterviewRouter.GET("classify")//面试题分类接口
		InterviewRouter.POST("list")//面试题列表接口
		InterviewRouter.GET("detail")//面试题详情接口
	}
	return InterviewRouter

}