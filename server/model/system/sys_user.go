package system

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户的UUID"`
	UserName string    `json:"user_name" gorm:"comment:用户名"`
	NickName string    `json:"nick_name" gorm:"comment:昵称"`
	Password string    `json:"-" gorm:"comment:密码"`
	Phone    string    `json:"phone" gorm:"comment:手机号"`
	Email    string    `json:"email" gorm:"comment:邮箱"`
	Avatar   string    `json:"avatar" gorm:"comment:头像"`
	RoleID   string    `json:"role_id" gorm:"comment:角色ID"`
	Role     string    `json:"role" gorm:"foreignKey:RoleID;references:ID;comment:角色"`
}
