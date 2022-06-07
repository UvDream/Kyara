package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
)

//SysApi 系统信息
type SysApi struct{}

// GetSystemApi 获取系统信息
func (s *SysApi) GetSystemApi(c *gin.Context) {
	if server, err := sysConfigService.GetSysConfig(); err != nil {
		global.Log.Error("获取系统信息失败(sysConfigService.GetSysConfig())", zap.String("err", err.Error()))
		response.FailWithMessage("获取系统信息失败", c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}
