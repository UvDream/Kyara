package v1

import (
	"server/global/response"
	"server/model"
	service "server/service/admin"
	"github.com/gin-gonic/gin"
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
