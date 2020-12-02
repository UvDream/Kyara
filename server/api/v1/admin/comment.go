package v1

import (
	"github.com/gin-gonic/gin"
	 resp "server/global/response"
	"server/model/request"
	"server/model/response"
	service "server/service/admin"
)

func GetBlogComment(c *gin.Context)  {
	var R request.ListStruct
	err := c.ShouldBindJSON(&R)
	if err!=nil{
		resp.FailWithMessage("获取参数错误",c)
	}
	err,list,msg,total:=service.GetBlogComment(R)
	if err != nil {
		resp.FailWithMessage(msg,c)

	}else{
		resp.OkDetailed(&response.PageResult{List: list,Total: total,Page: R.Page,PageSize: R.PageSize},msg,c)
	}
}