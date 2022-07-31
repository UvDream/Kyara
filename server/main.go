package main

import (
	"database/sql"
	"go.uber.org/zap"
	"server/core"
	"server/global"
	"server/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
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
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				global.Log.Error("关闭数据库链接失败", zap.Error(err))
			}
		}(db)
	}
	core.RunServer()
}
