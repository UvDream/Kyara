package service

import (
	"server/service/article"
	"server/service/file"
	"server/service/system"
)

type ServicesGroup struct {
	SystemServiceGroup  system.SysServiceGroup
	ArticleServiceGroup article.ArticlesServiceGroup
	FileServiceGroup    file.FilesServiceGroup
}

var ServicesGroupApp = new(ServicesGroup)
