package router

import (
	"server/router/article"
	"server/router/system"
)

type RoutersGroup struct {
	System  system.SysRouterGroup
	Article article.ArticlesGroup
}

var RoutersGroupApp = new(RoutersGroup)
