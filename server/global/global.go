package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/config"
)

var (
	Version   = "0.0.1"
	DB        *gorm.DB
	Viper     *viper.Viper      //读取配置文件
	Config    config.Config     //配置
	Log       *zap.Logger       //日志
	Redis     *redis.Client     //redis
	BlackList local_cache.Cache //jwt黑名单
)
