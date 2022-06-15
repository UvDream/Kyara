package request

import (
	"github.com/google/uuid"
	"server/model/common/request"
)

type ArticleRequest struct {
	UUID            uuid.UUID `json:"uuid" gorm:"comment:文章的UUID"`
	Title           string    `json:"title"  gorm:"type:varchar(100);not null" binding:"required"` // 标题
	Status          string    `json:"status" gorm:"type:varchar(10);not null" binding:"required"`  // 状态
	Slug            string    `json:"slug" gorm:"type:varchar(100);"`                              // 别名
	EditorType      string    `json:"editor_type" gorm:"type:varchar(100);"`                       // 编辑器类型
	MetaKeyWords    string    `json:"meta_key_words" gorm:"type:varchar(100);"`                    // 头部关键字
	MetaDescription string    `json:"meta_description" gorm:"type:varchar(100);"`                  // 头部描述
	Summary         string    `json:"summary" gorm:"type:varchar(100);"`                           // 摘要
	Thumbnail       string    `json:"thumbnail" gorm:"type:varchar(100);"`                         // 缩略图
	Visits          int       `json:"visits" gorm:"type:int(10);"`                                 // 访问量
	DisableComments bool      `json:"disable_comments" gorm:"type:tinyint(1);"`                    // 禁止评论
	Password        string    `json:"password" gorm:"type:varchar(100);"`                          // 访问密码
	Likes           int       `json:"likes" gorm:"type:int(10);"`                                  // 点赞数
	WordCount       int       `json:"word_count" gorm:"type:int(10);"`                             // 字数
	MdContent       string    `json:"md_content" gorm:"type:text;"`                                // markdown内容
	HtmlContent     string    `json:"html_content" gorm:"type:text;"`                              // html内容
	CommentCount    int       `json:"comment_count" gorm:"type:int(10);"`                          // 评论数
	TagsID          []uint    `json:"tags_id" gorm:"-"`                                            // tags id
	CategoryID      []uint    `json:"category_id" gorm:"-"`                                        // 分类id
	IsTop           bool      `json:"is_top" gorm:"type:tinyint(1);"`                              // 是否置顶
	AuthID          uuid.UUID `json:"auth_id" gorm:"comment:作者的UUID"`                              // 用户id
}

type ArticleListRequest struct {
	request.PaginationRequest
	Title      string `json:"title"` // 标题
	StartTime  string `form:"start_time" json:"start_time"`
	EndTime    string `form:"end_time" json:"end_time"`
	Status     string `form:"status" json:"status"`
	CategoryID uint   `form:"category_id" json:"category_id"`
	TagID      uint   `form:"tag_id" json:"tag_id"`
}