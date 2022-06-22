package article

import (
	"gorm.io/gorm"
	"server/model/article/request"
	"server/model/system"
)

type Article struct {
	gorm.Model
	request.ArticleRequest
	Tags       []Tag          `gorm:"many2many:tag_articles;foreignKey:UUID;joinForeignKey:ArticleID" json:"tags"`            // tags
	Categories []Category     `gorm:"many2many:category_articles;foreignKey:UUID;joinForeignKey:ArticleID" json:"categories"` // 分类
	Visits     int            `json:"visits" gorm:"type:int(10);"`                                                            // 访问量
	Likes      int            `json:"likes" gorm:"type:int(10);"`                                                             // 点赞数
	Auth       system.SysUser `gorm:"references:UUID" json:"author"`                                                          // 作者
}
