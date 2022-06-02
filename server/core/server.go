package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/initialize"
	"time"
)

func RunServer() {
	// 判断是否使用redis
	if global.Config.System.UseRedis || global.Config.System.UseMultipoint {
		initialize.Redis()
	}
	// 初始化路由器
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)
	global.Log.Info(s.ListenAndServe().Error())
}

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
