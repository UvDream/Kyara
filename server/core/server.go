package core

import (
	"server/global"
	"server/initialize"
)

func RunServer() {
	// 判断是否使用redis
	if global.Config.System.UseRedis || global.Config.System.UseMultipoint {
		initialize.Redis()
	}
	// 初始化路由器
	initialize.Routers()
}
