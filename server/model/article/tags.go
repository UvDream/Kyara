package article

import (
	"github.com/google/uuid"
	"server/model/common"
)

// Tag 标签
type Tag struct {
	common.Model
	Name      string `gorm:"type:varchar(100);" binding:"required" json:"name"` // 标签名
	Slug      string `gorm:"type:varchar(100);" json:"slug"`                    // 别名
	Thumbnail string `gorm:"type:varchar(100);" json:"thumbnail"`               // 缩略图
	Color     string `gorm:"type:varchar(100);" json:"color"`                   // 颜色
}

// TagArticle 标签和文章关联关系
type TagArticle struct {
	common.Model
	TagID     uint      `gorm:"type:int(10);not null"`                      // 标签ID
	ArticleID uuid.UUID `gorm:"type:varchar(100);not null;foreignKey:UUID"` // 文章ID
}
