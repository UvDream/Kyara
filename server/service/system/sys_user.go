package system

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/system"
	"server/utils"
)

type SysUserService struct{}

func (sysUserService *SysUserService) Login(username string, password string) (user *system.SysUser, token string, err error) {
	if global.DB != nil {
		return nil, "", fmt.Errorf("数据库未初始化")
	}
	if err := global.DB.Where("username = ?", username).Preload("Roles").First(&user).Error; err != nil {
		return nil, "", err
	} else {
		if ok := utils.BcryptCheck(password, user.Password); !ok {
			return nil, "", fmt.Errorf("密码错误")
		} else {
			//	颁发token
			j := &utils.JWT{
				SigningKey: []byte(global.Config.JWT.SigningKey),
			}
			claims := j.CreateClaims(request.BaseClaims{
				UUID:     user.UUID,
				ID:       user.ID,
				NickName: user.NickName,
				Username: user.UserName,
			})
			token, err := j.CreateToken(claims)
			if err != nil {
				global.Log.Error("颁发token失败", zap.Error(err))
				return nil, "", err
			} else {
				return user, token, nil
			}
		}
	}
}
