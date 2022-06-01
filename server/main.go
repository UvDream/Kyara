package main

import (
	"go.uber.org/zap"
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	// 初始化配置文件
	global.Viper = core.Viper()
	// 初始化日志
	global.Log = core.Zap()
	zap.ReplaceGlobals(global.Log)
	// 初始化数据库
	global.DB = initialize.Gorm()
	if global.DB != nil {
		// 初始化表
		initialize.RegistrationTable(global.DB)
		// 程序结束关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
}
