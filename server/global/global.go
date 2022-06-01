package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/config"
)

var (
	// Version is the current version of the application.
	Version = "0.0.1"
	DB      *gorm.DB
	Viper   *viper.Viper
	Config  config.Config
	Log     *zap.Logger
)
