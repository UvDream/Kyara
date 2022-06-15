package article

import (
	"gorm.io/gorm"
	"server/model/article/request"
)

type Article struct {
	gorm.Model
	request.ArticleRequest
	Tags     []Tag      `gorm:"-"`                           // tags
	Category []Category `gorm:"-"`                           // 分类
	Visits   int        `json:"visits" gorm:"type:int(10);"` // 访问量
	Likes    int        `json:"likes" gorm:"type:int(10);"`  // 点赞数
}
