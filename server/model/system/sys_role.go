package system

import (
	"server/model/common"
)

// SysRole 角色表
type SysRole struct {
	common.Model
	RoleName string     `json:"role_name" gorm:"not null;unique;primary_key;comment:角色名;size:90"` // 角色名
	ParentID string     `json:"parent_id" gorm:"comment:父角色ID"`                                   // 父角色ID
	Children []*SysRole `json:"children" gorm:"-"`
	// 关联下面的用户表
	Users []*SysUser `json:"-" gorm:"many2many:sys_user_role;"`
}
