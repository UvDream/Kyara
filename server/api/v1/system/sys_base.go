package system

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/global"
	"server/model/system"
	"server/utils"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
	})
}
func (b *BaseApi) Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func (b *BaseApi) InitData(c *gin.Context) {
	uid, _ := uuid.NewUUID()
	initializeUser := []system.SysUser{
		{
			UUID:     uid,
			UserName: "admin",
			Password: utils.BcryptHash("123456"),
			NickName: "admin",
			Avatar:   "www.pic.uvdream.cn",
			Phone:    "17621998888",
			Email:    "22222@163.com",
		},
	}
	global.DB.Create(&initializeUser)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "初始化数据成功",
	})
}
