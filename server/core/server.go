package core

import (
	"fmt"
	"server/global"
	"server/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	//InstallPlugs(Router)
	// end 插件描述

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Debug("server run success on ", address)

	fmt.Printf(`欢迎来到伽罗博客系统api服务
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)
	global.GVA_LOG.Error(s.ListenAndServe())
}
