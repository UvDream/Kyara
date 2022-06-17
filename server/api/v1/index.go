package v1

import (
	"server/api/v1/article"
	"server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiSystemGroup
	ArticleApiGroup  article.ApiArticleGroup
	TagApiGroup      article.TagsApi
	CategoryApiGroup article.CategoriesApi
}

var ApiGroupApp = new(ApiGroup)
