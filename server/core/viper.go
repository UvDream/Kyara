package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"server/global"
)

const (
	ConfigEnv         = "CONFIG"
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigDebugFile   = "config.debug.yaml"
	ConfigReleaseFile = "config.release.yaml"
)

// Viper 优先级:命令行 >配置文件 >默认值
func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "config", "config.yaml", "config file path")
		flag.Parse()
		//判断命令行参数是否为空
		if config == "" {
			if configEnv := os.Getenv(ConfigEnv); configEnv != "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = ConfigDebugFile
					fmt.Printf("你现在正在使用gin模式的%s环境,config的路径是%s\n", gin.EnvGinMode, ConfigDebugFile)
				case gin.TestMode:
					config = ConfigTestFile
					fmt.Printf("你现在正在使用gin模式的%s环境,config的路径是%s\n", gin.EnvGinMode, ConfigTestFile)
				case gin.ReleaseMode:
					config = ConfigReleaseFile
					fmt.Printf("你现在正在使用gin模式的%s环境,config的路径是%s\n", gin.EnvGinMode, ConfigReleaseFile)
				default:
					config = ConfigDefaultFile
					fmt.Printf("你现在正在使用gin模式的%s环境,config的路径是%s\n", gin.EnvGinMode, ConfigDefaultFile)
				}
			} else {
				config = configEnv
				fmt.Printf("你现在正在使用gin模式的%s环境,config的路径是%s\n", ConfigEnv, config)
			}
		} else {
			//	命令行参数不为空
			fmt.Printf("你正在使用命令行的-c 参数传递的值,config的路径为%s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("你正在使用Viper函数传递的值,config的路径为%s\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取错误: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生变化:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			panic(fmt.Errorf("配置文件反序列化错误: %s \n", err))
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("配置文件反序列化错误: %s \n", err))
	}
	return v
}
