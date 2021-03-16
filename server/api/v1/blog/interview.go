package blog

import (
	"github.com/gin-gonic/gin"
	"server/global/response"
	"server/model"
	"server/model/request"
	"server/service/blog"
	"server/utils"
	resp "server/model/response"
)
//新增面试题分类
func AddInterviewClassify(c *gin.Context)  {
	var R  model.InterviewClassify
	err:=c.ShouldBindJSON(&R)
	if err!=nil {
		response.FailWithMessage("获取参数错误",c)
	}
	apiData:=utils.Rules{
		"ClassifyName":{utils.NotEmpty()},
		"ClassifyIcon": {utils.NotEmpty()},
	}
	err=utils.Verify(R,apiData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err,msg:=blog.AddInterviewClassify(R)
	if err!=nil{
		response.FailWithMessage(msg,c)
	}else{
		response.OkWithMessage(msg,c)
	}

}
//删除面试题分类
func DeleteInterviewClassify(c *gin.Context)  {
	id:=c.Query("id")
	if id=="" {
		response.FailWithMessage("缺少分类id", c)
	}
	err,msg:=blog.DeleteInterviewClassify(id)
	if err!=nil{
		response.FailWithMessage(msg,c)
	}else{
		response.OkWithMessage(msg,c)
	}

}
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
	}
	err=utils.Verify(R,apiData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err,msg,data,totalCount:=blog.GetInterview(R)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(resp.InterviewResponse{Msg:msg,Data:data,TotalCount:totalCount}, msg, c)
	}
}
//获取面试题详情
func GetInterviewDetail(c *gin.Context)  {
	 id:=c.Query("id")
	if id=="" {
		response.FailWithMessage("缺少id", c)
	}
	err,msg,data:=blog.GetInterviewDetail(id)
	if err!=nil{
		response.FailWithMessage(msg,c)
	}else{
		response.OkDetailed(data,msg,c)
	}
}

//新增面试题
func AddInterview(c *gin.Context)  {
	var R model.Interview
	err:=c.ShouldBindJSON(&R)
	if err!=nil {
		response.FailWithMessage("获取参数错误",c)
		return
	}
	apiData:=utils.Rules{
		"Title":{utils.NotEmpty()},
		"ClassifyID":{utils.NotEmpty()},
	}
	err=utils.Verify(R,apiData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err,msg,data:=blog.AddInterview(R)
	if err!=nil{
		response.FailWithMessage(msg,c)
	}else{
		response.OkDetailed(data,msg,c)
	}

}