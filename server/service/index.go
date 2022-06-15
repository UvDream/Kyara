package service

import (
	"server/service/article"
	"server/service/system"
)

type ServicesGroup struct {
	SystemServiceGroup  system.SysServiceGroup
	ArticleServiceGroup article.ArticlesServiceGroup
}

var ServicesGroupApp = new(ServicesGroup)
