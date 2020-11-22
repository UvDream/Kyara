package model

import "github.com/jinzhu/gorm"

// 博客留言
type  BlogComment struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"comment:'昵称'"`
	Avatar string `json:"avatar" gorm:"comment:'用户头像'"`
	Email string `json:"email"`
	BlogURL string `json:"blog_url"`
	CommentContent string `json:"comment_content" gorm:"comment:'留言内容'"`
	UserID string `json:"user_id" gorm:"comment:'登陆用户用户id'"`
	LikeNumber string `json:"like_number" gorm:"comment:'点赞数'"`
	BadNumber string `json:"bad_number" gorm:"comment:'差评数'"`
	ParentID string `json:"parent_id" gorm:"comment:'父id'"`
	IsPrivate string `json:"is_private" gorm:"comment:'是否为私密评论'"`
	Status string `json:"status" gorm:"comment:'审核状态0:需要审核,1:可以展示'`
}