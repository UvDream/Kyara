package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"server/global"
	"server/initialize/internal"
	"server/model/system"
)

func GormMysql() *gorm.DB {
	m := global.Config.Mysql
	// 无数据库名称
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func RegistrationTable(db *gorm.DB) {
	//初始化表注册
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysRole{},
		system.SysUserRole{},
	)
	if err != nil {
		global.Log.Error("注册数据表失败", zap.Error(err))
		os.Exit(0)
	}
	global.Log.Info("注册数据表成功")
}
