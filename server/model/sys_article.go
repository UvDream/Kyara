package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

// SysArticle 文章数据表
type SysArticle struct {
	gorm.Model
	//更新时间
	UpdateTime  time.Time `json:"update_time" gorm:"comment:'更新时间'"`
	//文章id
	ArticleID uint `json:"article_id" gorm:"-"`
	// 文章标题
	Title string `json:"title" gorm:"comment:'文章标题'"`
	// 文章摘要
	Introduction string `json:"introduction" gorm:"comment:'文章摘要'"`
	// 作者id
	UserID uuid.UUID `json:"user_id" gorm:"comment:'作者id'"`
	// ---作者名---
	UserName string `json:"user_name" gorm:"-"`
	// 评论id
	CommentID string `json:"comment_id" gorm:"comment:'评论id'"`
	// icon
	Icon string `json:"icon" gorm:"comment:'icon';default:'uv-newspaper-outline'"`
	// icon颜色
	IconColor string `json:"icon_color" gorm:"comment:'icon颜色'"`
	// 图片地址
	ImgURL string `json:"img_url" gorm:"comment:'图片地址'"`
	// 文章内容
	ArticleContent string `json:"article_content" gorm:"comment:'文章内容';type:longblob"`
	//---------html内容
	ArticleHtml string `json:"article_html" gorm:"-"`
	// 浏览量
	ViewCount string `json:"view_count" gorm:"comment:'浏览量'"`
	// 原文链接
	OriginalLink string `json:"original_link" gorm:"comment:'原文链接'"`
	// 标签
	TagArray []uint `json:"tag_array" gorm:"-"`
	// 分类id
	ClassificationID string `json:"classification_id" gorm:"comment:'分类id'"`
	// 阅读密码
	ViewPassword string `json:"view_password" gorm:"comment:'阅读密码'"`
	// 是否开启评论
	IsComment string `json:"is_comment" gorm:"comment:'是否开启评论'"`
	// 评论量
	CommentCount string `json:"comment_count" gorm:"comment:'评论量'"`
	// 状态
	Status string `json:"status" gorm:"comment:'状态';default:'0'"`
	// 是否需要密码
	IsPassword string `json:"is_password" gorm:"comment:'是否需要密码';default:'0'"`
	// 转载规则设置
	TransferRules string `json:"transfer_rules" gorm:"comment:'转载规则设置';default:'0'"`
	// 封面类型
	CoverType string `json:"cover_type" gorm:"comment:'封面类型';default:'0'"`
	// 是否置顶
	Top string `json:"top" gorm:"comment:'是否置顶';default:'0'"`
	// 字数
	WordCount uint `json:"word_count" gorm:"comment:'文章字数';default:'0'"`
	// 点赞数
	LikeCount string `json:"like_count" gorm:"comment:'点赞数'"`
	// 赞赏码
	CollectList []CollectionCode `json:"collect_list" gorm:"-"`
	CollectText string           `json:"collect_text" gorm:"comment:'赞赏提示文字';default:'如果觉得我的文章对你有用，请随意赞赏!'"`
}

//文章分类
type SysClassify struct {
	gorm.Model
	TypeName  string        `json:"type_name" gorm:"comment:'分类名称'"`
	Icon      string        `json:"icon"`
	ParentID  string        `json:"parent_id"`
	Children  []SysClassify `json:"children"`
	IconColor string        `json:"icon_color" gorm:"comment:'icon颜色';default:'#2d8cf0'"`
}

// 文章标签
type SysTag struct {
	gorm.Model
	TagName  string `json:"tag_name"`
	TagCount int64  `json:"tag_count" gorm:"-"`
	Color    string `json:"color" gorm:"comment:'颜色';default:'#108ee9'"`
}

//点赞表
type SysLike struct {
	gorm.Model
	UserID    string `json:"user_id" gorm:"comment:'点赞用户'"`
	ArticleID string `json:"article_id" gorm:"comment:'文章id'"`
}

//文章以及tag关联表
type SysArticleTag struct {
	gorm.Model
	ArticleID uint `json:"article_id"`
	TagID     uint `json:"tag_id"`
}

//ip地址
type BlogView struct {
	gorm.Model
	IP     string `json:"ip"`
	System string `json:"system" gorm:"comment:'系统'"`
	Device string `json:"device" gorm:"comment:'设备类型'"`
}

//收款码
type CollectionCode struct {
	gorm.Model
	Name      string `json:"name" gorm:"comment:'收款码名称'"`
	ImgURL    string `json:"img_url" gorm:"comment:'收款码地址'"`
	ArticleID uint `json:"article_id" gorm:"comment:'对应的文章id'"`
	UserID    uuid.UUID `json:"user_id" gorm:"comment:'上传作者id';default:'ce0d6685-c15f-4126-a5b4-890bc9d2356d'"`
	Type string `json:"type" gorm:"comment:'默认赞赏码为0'"`
}
