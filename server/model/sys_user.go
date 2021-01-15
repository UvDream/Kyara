package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:'用户UUID'"`
	Username    string       `json:"userName" gorm:"comment:'用户登录名'"`
	Password    string       `json:"-"  gorm:"comment:'用户登录密码'"`
	NickName    string       `json:"nickName" gorm:"default:'系统用户';comment:'用户昵称'" `
	HeaderImg   string       `json:"headerImg" gorm:"default:'https://gitee.com/Uvdream/images/raw/master/images/20200825150341.png';comment:'用户头像'"`
	Authority   SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId;comment:'用户角色'"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:'用户角色ID'"`
	BlogURL     string       `json:"blog_url" gorm:"comment:'博客地址'"`
	Email       string       `json:"email" gorm:"comment:'邮箱'"`
}
