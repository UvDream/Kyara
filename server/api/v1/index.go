package v1

import (
	"server/api/v1/article"
	"server/api/v1/file"
	"server/api/v1/system"
	"server/api/v1/theme"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiSystemGroup
	ArticleApiGroup  article.ApiArticleGroup
	TagApiGroup      article.TagsApi
	CategoryApiGroup article.CategoriesApi
	FileApiGroup     file.ApiFileGroup
	ThemeApiGroup    theme.ApiThemeGroup
}

var ApiGroupApp = new(ApiGroup)
