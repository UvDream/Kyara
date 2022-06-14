package system

import "server/service"

type ApiSystemGroup struct {
	UserApi
	BaseApi
	SysApi
}

var (
	sysConfigService = service.ServicesGroupApp.SystemServiceGroup.SysConfigService
	userService      = service.ServicesGroupApp.SystemServiceGroup.SysUserService
)
