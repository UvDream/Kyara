package v1

import (
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	service "gin-vue-admin/service/admin"
	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var R model.SysArticle
	err := c.ShouldBind(&R)
	if err != nil {
		//	参数获取失败
	}
	err, msg := service.AddArticle(R)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkWithMessage(msg, c)
	}
}
