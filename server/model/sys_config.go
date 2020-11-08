package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// SysConfig 博客配置
type SysConfig struct {
	gorm.Model
	//博客相关设置
	BlogName string `json:"blog_name" gorm:"comment:'博客名称';default:'U世界的V梦想'"`
	BlogLogo string `json:"blog_logo" gorm:"comment:'博客logo';default:'https://pic.baixiongz.com/2020/08/05/a46e139ec4236.png'"`
	BlogNotice string `json:"blog_notice" gorm:"comment:'博客公告'"`
	BlogStartTime time.Time `json:"blog_start_time" gorm:"comment:'博客开始时间';default:'2015-01-02 15:04:05'"`
	BlogViewCount string `json:"blog_view_count" gorm:"comment:'博客访问量'"`
	// 作者相关设置
	AuthorAvatar string `json:"author_avatar" gorm:"comment:'作者头像';default:'https://pic.baixiongz.com/2020/08/05/a46e139ec4236.png'"`
	AuthorName string `json:"author_name" gorm:"comment:'作者名称';default:'UvDream'"`
	AuthorLink string `json:"author_link" gorm:"comment:'点击作者头像跳转地址';default:'https://github.com/UvDream'"`
	AuthorMotto string `json:"author_motto" gorm:"comment:'作者座右铭';default:'一切皆有可能!'"`
	ActiveTime time.Time `json:"active_time" gorm:"comment:'最近活动时间';default:'2020-09-12 15:04:05'"`
	// 图床相关设置
	ImgurUserName string `json:"imgur_user_name" gorm:"comment:'图床用户名'"`
	ImgurPassword string `json:"imgur_password" gorm:"comment:'图床密码'"`
	ImgurToken string `json:"imgur_token" gorm:"comment:'图床token'"`
	ImgurURL string `json:"imgur_url" gorm:"comment:'图床连接'"`
	ImgurType string `json:"imgur_type" gorm:"comment:'图床类型';default:'0'"`
	// github
	GithubUserName string `json:"github_user_name" gorm:"comment:'github用户名'"`
}
//打赏库
type Rewards struct {
	gorm.Model
	RewardName string `json:"reward_name" gorm:"comment:'打赏名称'"`
	RewardImgURL string `json:"reward_img_url" gorm:"comment:'打赏码连接'"`
	UserID string `json:"user_id" gorm:"comment:'用户id'"`
}
