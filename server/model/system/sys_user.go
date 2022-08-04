package system

import (
	"server/model/common"
)

// SysUser 用户表
type SysUser struct {
	common.Model
	UserName string `json:"user_name" gorm:"comment:用户名" binding:"required"`
	NickName string `json:"nick_name" gorm:"comment:昵称"`
	Password string `json:"password" gorm:"comment:密码" binding:"required"`
	Phone    string `json:"phone" gorm:"comment:手机号"`
	Email    string `json:"email" gorm:"comment:邮箱"`
	Avatar   string `json:"avatar" gorm:"comment:头像"`
	//关联到角色表
	Roles []SysRole `json:"roles" gorm:"many2many:sys_user_role;"`
}
