package system

import "server/service"

type ApiGroup struct {
	UserApi
	BaseApi
	SysApi
}

var (
	sysConfigService = service.ServicesGroupApp.SystemServiceGroup.SysConfigService
	userService      = service.ServicesGroupApp.SystemServiceGroup.SysUserService
)
