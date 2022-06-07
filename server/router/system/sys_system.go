package system

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type SysRouter struct{}

func (s *SysRouter) InitSysRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system")
	systemApi := v1.ApiGroupApp.SystemApiGroup.SysApi
	{
		sysRouter.GET("get_system_config", systemApi.GetSystemApi)
	}
}
