package router

import "server/router/system"

type RoutersGroup struct {
	System system.SysRouterGroup
}

var RoutersGroupApp = new(RoutersGroup)
