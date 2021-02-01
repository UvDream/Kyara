package v1

import (
	"github.com/gin-gonic/gin"
	"server/global/response"
	"server/model"
	service "server/service/admin"
)

func SaveConfig(c *gin.Context)  {
	var R model.SysConfig
	err := c.ShouldBindJSON(&R)
	if err!=nil {
		response.FailWithMessage("获取参数失败",c)
	}
	err,msg:=service.SaveConfig(R)
	if err!=nil {
		response.FailWithMessage(msg,c)
	}else{
		response.OkWithMessage(msg,c)
	}
}
