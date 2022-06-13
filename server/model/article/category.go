package article

import "gorm.io/gorm"

//Category 分类
type Category struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"` // 分类名
	Slug        string `gorm:"type:varchar(100);not null"` // 别名
	Description string `gorm:"type:varchar(100);not null"` // 描述
	Thumbnail   string `gorm:"type:varchar(100);not null"` // 缩略图
	ParentID    int    `gorm:"type:int(10);not null"`      // 父级ID
	Password    string `gorm:"type:varchar(100);not null"` // 访问密码
}

// CategoryArticle 分类和文章关联关系
type CategoryArticle struct {
	gorm.Model
	CategoryID uint `gorm:"type:int(10);not null"` // 分类ID
	ArticleID  uint `gorm:"type:int(10);not null"` // 文章ID
}
