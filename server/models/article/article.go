package article

import (
	"server/models"
	"server/models/article/request"
	"server/models/system"
)

type Article struct {
	models.Model
	request.ArticleRequest
	Tags       []Tag       `gorm:"many2many:tag_articles;joinForeignKey:ArticleID" json:"tags"`            // tags
	Categories []Category  `gorm:"many2many:category_articles;joinForeignKey:ArticleID" json:"categories"` // 分类
	Visits     int         `json:"visits" gorm:"type:int(10);"`                                            // 访问量
	Likes      int         `json:"likes" gorm:"type:int(10);"`                                             // 点赞数
	Auth       system.User `json:"author"`                                                                 // 作者
}
