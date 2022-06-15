package article

import "server/service"

type ApiArticleGroup struct {
	ArticlesApi
}

var (
	articleService = service.ServicesGroupApp.ArticleServiceGroup.ToArticleService
)
