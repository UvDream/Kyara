package router

import (
	"server/router/article"
	"server/router/system"
)

type RoutersGroup struct {
	System   system.SysRouterGroup
	Article  article.ArticlesGroup
	Tag      article.TagsStruct
	Category article.CategoriesStruct
}

var RoutersGroupApp = new(RoutersGroup)
