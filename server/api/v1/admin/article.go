package v1

import (
	"github.com/gin-gonic/gin"
	"server/global/response"
	"server/model"
	service "server/service/admin"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var R model.SysArticle
	err := c.ShouldBindJSON(&R)
	if err != nil {
		//	参数获取失败
	}
	err, msg ,data:= service.AddArticle(R,c)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(data,msg, c)
	}
}
//查询文章详情
func GetArticleDetail(c *gin.Context)  {
	id := c.Query("id")
	err,data:=service.GetArticleDetail(id)
	if err != nil {
		response.FailWithMessage("获取详情错误", c)
	} else {
		response.OkDetailed(data,"获取详情成功", c)
	}
}
