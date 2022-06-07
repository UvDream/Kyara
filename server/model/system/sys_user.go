package system

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SysUser 用户表
type SysUser struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户的UUID"`
	UserName string    `json:"user_name" gorm:"comment:用户名"`
	NickName string    `json:"nick_name" gorm:"comment:昵称"`
	Password string    `json:"-" gorm:"comment:密码"`
	Phone    string    `json:"phone" gorm:"comment:手机号"`
	Email    string    `json:"email" gorm:"comment:邮箱"`
	Avatar   string    `json:"avatar" gorm:"comment:头像"`
	//关联到角色表
	Roles []SysRole `json:"roles" gorm:"many2many:sys_user_role;"`
}
