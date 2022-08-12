package system

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	"server/models"
	"server/models/system"
	"server/utils"
)

type SysUserService struct{}

func (sysUserService *SysUserService) Login(username string, password string) (user *system.User, token string, msg string, err error) {
	//先判定数据库是否存在
	if global.DB == nil {
		return nil, "", "数据库不存在", fmt.Errorf("数据库未初始化")
	}
	if err := global.DB.Where("user_name = ?", username).Preload("Roles").First(&user).Error; err != nil {
		return nil, "", "查找用户信息错误", err
	} else {
		if ok := utils.BcryptCheck(password, user.Password); !ok {
			return nil, "", "密码错误", fmt.Errorf("密码错误")
		} else {
			//	颁发token
			j := &utils.JWT{
				SigningKey: []byte(global.Config.JWT.SigningKey),
			}
			claims := j.CreateClaims(models.BaseClaims{
				ID:       user.ID,
				NickName: user.NickName,
				Username: user.UserName,
			})
			token, err := j.CreateToken(claims)
			if err != nil {
				global.Log.Error("颁发token失败", zap.Error(err))
				return nil, "", "token生成失败", err
			} else {
				return user, token, "登陆成功", nil
			}
		}
	}
}
