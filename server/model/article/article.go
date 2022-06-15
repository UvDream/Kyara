package article

import (
	"gorm.io/gorm"
	"server/model/article/request"
)

type Article struct {
	gorm.Model
	request.ArticleRequest
	Tags     []Tag      `gorm:"-"` // tags
	Category []Category `gorm:"-"` // 分类
}
