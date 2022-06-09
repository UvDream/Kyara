package system

import (
	"fmt"
	"server/global"
	"server/model/system"
	"server/utils"
)

type SysUserService struct{}

func (sysUserService *SysUserService) Login(username string, password string) (userInfo *system.SysUser, err error) {
	if global.DB != nil {
		return nil, fmt.Errorf("数据库未初始化")
	}
	var user system.SysUser
	if err := global.DB.Where("username = ?", username).Preload("Roles").First(&user).Error; err != nil {
		return nil, err
	} else {
		if ok := utils.BcryptCheck(password, user.Password); !ok {
			return nil, fmt.Errorf("密码错误")
		} else {
			//	颁发token

		}
	}
	return
}
