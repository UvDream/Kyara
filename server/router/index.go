package router

import (
	"server/router/article"
	"server/router/file"
	"server/router/system"
	"server/router/theme"
)

type RoutersGroup struct {
	System   system.SysRouterGroup
	Article  article.ArticlesGroup
	Tag      article.TagsStruct
	Category article.CategoriesStruct
	File     file.FilesRouterGroup
	Theme    theme.ThemesGroupRouter
}

var RoutersGroupApp = new(RoutersGroup)
