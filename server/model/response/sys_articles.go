package response

import (
	"server/model"
	"time"
)

type SysArticleListResponse struct {
	Msg        []model.SysArticle `json:"msg"`
	TotalCount int64              `json:"totalCount"`
}
type SysConfigsResponse struct {
	//博客相关设置
	BlogName      string    `json:"blog_name" gorm:"comment:'博客名称';default:'U世界的V梦想'"`
	BlogLogo      string    `json:"blog_logo" gorm:"comment:'博客logo';default:'https://pic.baixiongz.com/2020/08/05/a46e139ec4236.png'"`
	BlogNotice    string    `json:"blog_notice" gorm:"comment:'博客公告'"`
	BlogStartTime time.Time `json:"blog_start_time" gorm:"comment:'博客开始时间';default:'2015-01-02 15:04:05'"`
	BlogViewCount string    `json:"blog_view_count" gorm:"comment:'博客访问量'"`
	// 作者相关设置
	AuthorAvatar string    `json:"author_avatar" gorm:"comment:'作者头像';default:'https://pic.baixiongz.com/2020/08/05/a46e139ec4236.png'"`
	AuthorName   string    `json:"author_name" gorm:"comment:'作者名称';default:'UvDream'"`
	AuthorLink   string    `json:"author_link" gorm:"comment:'点击作者头像跳转地址';default:'https://github.com/UvDream'"`
	AuthorMotto  string    `json:"author_motto" gorm:"comment:'作者座右铭';default:'一切皆有可能!'"`
	ActiveTime   time.Time `json:"active_time" gorm:"comment:'最近活动时间';default:'2020-09-12 15:04:05'"`
	// 文章
	ArticleCount string `json:"article_count"`
}

//github Module
type GithubList struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	NodeID      string `json:"node_id"`
	Description string `json:"description"`
	//fork数量
	ForksCount string `json:"forks_count"`
	//	star数量
	StargazersCount string `json:"stargazers_count"`
	//地址
	HtmlURL string `json:"html_url"`
	//	语言
	Language string `json:"language"`
}
