package v1

import (
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(c *gin.Context)  {
	var R model.SysArticle
	err:=c.ShouldBind(&R)
	if err!=nil {
	//	参数获取失败
	}
}
