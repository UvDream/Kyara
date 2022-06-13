package article

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UUID            uuid.UUID  `json:"uuid" gorm:"comment:文章的UUID"`
	Title           string     `gorm:"type:varchar(100);not null"` // 标题
	Status          int        `gorm:"type:tinyint(1);not null"`   // 状态
	Slug            string     `gorm:"type:varchar(100);not null"` // 别名
	EditorType      string     `gorm:"type:varchar(100);not null"` // 编辑器类型
	MetaKeyWords    string     `gorm:"type:varchar(100);not null"` // 头部关键字
	MetaDescription string     `gorm:"type:varchar(100);not null"` // 头部描述
	Summary         string     `gorm:"type:varchar(100);not null"` // 摘要
	Thumbnail       string     `gorm:"type:varchar(100);not null"` // 缩略图
	Visits          int        `gorm:"type:int(10);not null"`      // 访问量
	DisableComments bool       `gorm:"type:tinyint(1);not null"`   // 禁止评论
	Password        string     `gorm:"type:varchar(100);not null"` // 访问密码
	Likes           int        `gorm:"type:int(10);not null"`      // 点赞数
	WordCount       int        `gorm:"type:int(10);not null"`      // 字数
	MdContent       string     `gorm:"type:text;not null"`         // markdown内容
	HtmlContent     string     `gorm:"type:text;not null"`         // html内容
	CommentCount    int        `gorm:"type:int(10);not null"`      // 评论数
	Tags            []Tag      `gorm:"-"`                          // tags
	TagsID          []uint     `gorm:"-"`                          // tags id
	Category        []Category `gorm:"-"`                          // 分类
	CategoryID      []uint     `gorm:"-"`                          // 分类id
	IsTop           bool       `gorm:"type:tinyint(1);not null"`   // 是否置顶
}
