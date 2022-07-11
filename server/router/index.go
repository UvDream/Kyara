package router

import (
	"server/router/article"
	"server/router/file"
	"server/router/system"
)

type RoutersGroup struct {
	System   system.SysRouterGroup
	Article  article.ArticlesGroup
	Tag      article.TagsStruct
	Category article.CategoriesStruct
	File     file.FilesRouterGroup
}

var RoutersGroupApp = new(RoutersGroup)
