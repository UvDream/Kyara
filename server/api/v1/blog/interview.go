package blog

import (
	"github.com/gin-gonic/gin"
	"server/global/response"
	"server/model/request"
	"server/service/blog"
	"server/utils"
	resp "server/model/response"
)
//获取面试题分类
func GetInterviewClassify(c *gin.Context)  {
	err,msg,data:=blog.GetInterviewClassify()
	if err!=nil{
		 response.FailWithMessage(msg,c)
	}else{
		response.OkDetailed(data,msg,c)
	}
}

//获取面试题list
func GetInterview(c *gin.Context)  {
	var R request.InterviewSearch
	err:=c.ShouldBindJSON(&R)
	if err!=nil {
		response.FailWithMessage("获取参数错误",c)
	}
	apiData:=utils.Rules{
		"Page":{utils.NotEmpty()},
		"PageSize":{utils.NotEmpty()},
		"Classify":{utils.NotEmpty()},
	}
	err=utils.Verify(R,apiData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err,msg,data,totalCount:=blog.GetInterview(c,R)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(resp.InterviewResponse{Msg:msg,Data:data,TotalCount:totalCount}, msg, c)
	}
}
