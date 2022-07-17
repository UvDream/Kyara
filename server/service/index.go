package service

import (
	"server/service/article"
	"server/service/file"
	"server/service/system"
	"server/service/theme"
)

type ServicesGroup struct {
	SystemServiceGroup  system.SysServiceGroup
	ArticleServiceGroup article.ArticlesServiceGroup
	FileServiceGroup    file.FilesServiceGroup
	ThemeServiceGroup   theme.ThemesServiceGroup
}

var ServicesGroupApp = new(ServicesGroup)
