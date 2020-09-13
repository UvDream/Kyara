package model
import (
	"github.com/jinzhu/gorm"
)
// SysComment 评论model
type SysComment struct{
	gorm.Model
	CommentContent string `json:"comment_content" gorm:"comment:'评论内容'"`
	UserID string `json:"user_id" gorm:"comment:'作者id'"`
	ParentID string `json:"parent_id" gorm:"comment:'父节点';default:''"`
	IsPrivate string `json:"is_private" gorm:"comment:'是否为私密评论'"`
	ArticleID string `json:"article_id" gorm:"comment:'文章id,文章id可查出作者'"`
}
