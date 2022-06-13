package article

import "gorm.io/gorm"

// Tag 标签
type Tag struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);not null"` // 标签名
	Slug      string `gorm:"type:varchar(100);not null"` // 别名
	Thumbnail string `gorm:"type:varchar(100);not null"` // 缩略图
	Color     string `gorm:"type:varchar(100);not null"` // 颜色
}

// TagArticle 标签和文章关联关系
type TagArticle struct {
	gorm.Model
	TagID     uint `gorm:"type:int(10);not null"` // 标签ID
	ArticleID uint `gorm:"type:int(10);not null"` // 文章ID
}
