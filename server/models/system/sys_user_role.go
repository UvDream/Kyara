package system

// SysUserRole 用户和角色关联表
type SysUserRole struct {
	SysUserId uint   `gorm:"column:sys_user_id"`
	SysRoleId string `gorm:"column:sys_role_id"`
}
