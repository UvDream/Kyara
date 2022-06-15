package request

import "github.com/google/uuid"

type ArticleRequest struct {
	UUID            uuid.UUID `json:"uuid" gorm:"comment:文章的UUID"`
	Title           string    `json:"title"  gorm:"type:varchar(100);not null" binding:"required"` // 标题
	Status          int       `json:"status"  gorm:"type:tinyint(1);not null"`                     // 状态
	Slug            string    `json:"slug" gorm:"type:varchar(100);not null"`                      // 别名
	EditorType      string    `json:"editor_type" gorm:"type:varchar(100);not null"`               // 编辑器类型
	MetaKeyWords    string    `json:"meta_key_words" gorm:"type:varchar(100);not null"`            // 头部关键字
	MetaDescription string    `json:"meta_description" gorm:"type:varchar(100);not null"`          // 头部描述
	Summary         string    `json:"summary" gorm:"type:varchar(100);not null"`                   // 摘要
	Thumbnail       string    `json:"thumbnail" gorm:"type:varchar(100);not null"`                 // 缩略图
	Visits          int       `json:"visits" gorm:"type:int(10);not null"`                         // 访问量
	DisableComments bool      `json:"disable_comments" gorm:"type:tinyint(1);not null"`            // 禁止评论
	Password        string    `json:"password" gorm:"type:varchar(100);not null"`                  // 访问密码
	Likes           int       `json:"likes" gorm:"type:int(10);not null"`                          // 点赞数
	WordCount       int       `json:"word_count" gorm:"type:int(10);not null"`                     // 字数
	MdContent       string    `json:"md_content" gorm:"type:text;not null"`                        // markdown内容
	HtmlContent     string    `json:"html_content" gorm:"type:text;not null"`                      // html内容
	CommentCount    int       `json:"comment_count" gorm:"type:int(10);not null"`                  // 评论数
	TagsID          []uint    `json:"tags_id" gorm:"-"`                                            // tags id
	CategoryID      []uint    `json:"category_id" gorm:"-"`                                        // 分类id
	IsTop           bool      `json:"is_top" gorm:"type:tinyint(1);not null"`                      // 是否置顶
	AuthID          uuid.UUID `json:"auth_id" gorm:"type:int(10);not null"`                        // 用户id
}
