package system

import (
	"github.com/google/uuid"
	code2 "server/code"
	"server/global"
	"server/model/system"
	"server/utils"
)

// RegisterService 注册用户
func (*SysUserService) RegisterService(opts system.SysUser) (user system.SysUser, code int, err error) {
	db := global.DB
	//首先查询账号是否存在
	if err := db.Where("user_name = ?", opts.UserName).First(&user).Error; err == nil {
		return opts, code2.ErrorUserExist, err
	}
	//查询昵称是否存在
	if err := db.Where("nick_name = ?", opts.NickName).First(&user).Error; err == nil {
		return opts, code2.ErrorUserExist, err
	}
	//查询邮箱是否存在
	if err := db.Where("email = ?", opts.Email).First(&user).Error; err == nil {
		return opts, code2.ErrorUserExistEmail, err
	}
	//查询手机号是否存在
	if err := db.Where("phone = ?", opts.Phone).First(&user).Error; err == nil {
		return opts, code2.ErrorUserExistPhone, err
	}
	opts.UUID = uuid.New()
	opts.Password = utils.BcryptHash(opts.Password)
	if err := db.Create(&opts).Error; err != nil {
		return opts, code2.ErrorCreateUser, err
	}
	return opts, code2.SUCCESS, err
}
