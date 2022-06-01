package initialize

import (
	"gorm.io/gorm"
	"server/global"
)

func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgsql()
	case "sqlite3":
		return GormSqlite3()
	default:
		return GormMysql()
	}
}
