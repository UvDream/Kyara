package article

import "server/service"

type ApiArticleGroup struct {
	ArticlesApi
	CategoriesApi
	TagsApi
}

var (
	articleService = service.ServicesGroupApp.ArticleServiceGroup.ToArticleService
	tagService     = service.ServicesGroupApp.ArticleServiceGroup.ToTagService
)
