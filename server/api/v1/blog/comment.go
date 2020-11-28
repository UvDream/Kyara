package blog

import (
	"github.com/gin-gonic/gin"
	"server/global/response"
	resp "server/model/response"
	"server/model"
	"server/model/request"
	"server/service/blog"
	"server/utils"
)

//博客留言
func Comment(c *gin.Context) {
	var R model.BlogComment
	_ = c.ShouldBindJSON(&R)
	if R.ID==0{
		ApiVerify := utils.Rules{
			"UserName":{utils.NotEmpty()},
			"Email":{utils.NotEmpty()},
			"Avatar":{utils.NotEmpty()},
			"CommentContent": {utils.NotEmpty()},
		}
		ApiVerifyErr := utils.Verify(R, ApiVerify)
		if ApiVerifyErr != nil {
			response.FailWithMessage(ApiVerifyErr.Error(), c)
			return
		}
	}

	err,msg:=blog.Comment(R)
	if err!=nil{
		response.FailWithMessage(msg,c)
	}else {
		response.OkWithMessage(msg,c)
	}
}
//获取留言
func GetComment(c *gin.Context)  {
	var R request.ListStruct
	_ = c.ShouldBindJSON(&R)
	err,msg,list,count:=blog.GetComment(R)
	if err!=nil{
		response.FailWithMessage(msg,c)
	}else {
		response.OkDetailed(&resp.CommentMsg{Total: count,Data: list},msg,c)
	}
}
//百度收录
func ToBaiDu(c *gin.Context)  {
	var R request.ToBaiDuRequest
	err:=c.ShouldBindJSON(&R)
	if err!=nil {
		response.FailWithMessage("参数获取错误",c)
	}
	err,msg,count:=blog.ToBaidu(R)
	if err!=nil{
		response.FailWithMessage(msg,c)
	}else {
		response.OkDetailed(count,msg,c)
	}
}
