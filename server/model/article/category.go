package article

import (
	"server/model/common"
)

//Category 分类
type Category struct {
	common.Model
	Name        string     `gorm:"type:varchar(100);" binding:"required" json:"name"` // 分类名
	Slug        string     `gorm:"type:varchar(100);" json:"slug"`                    // 别名
	Description string     `gorm:"type:varchar(1000);" json:"description"`            // 描述
	Thumbnail   string     `gorm:"type:varchar(100);" json:"thumbnail"`               // 缩略图
	ParentID    int        `gorm:"type:int(10);" json:"parent_id"`                    // 父级ID
	Password    string     `gorm:"type:varchar(100);" json:"password"`                // 访问密码
	Children    []Category `gorm:"-" json:"children"`                                 //子类
}

// CategoryArticle 分类和文章关联关系
type CategoryArticle struct {
	common.Model
	CategoryID uint   `gorm:"type:int(10);not null"`                 // 分类ID
	ArticleID  string `gorm:"varchar(100);not null;foreignKey:UUID"` // 文章ID
}
