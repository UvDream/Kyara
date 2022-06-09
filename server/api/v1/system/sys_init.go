package system

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/global"
	"server/model/system"
	"server/utils"
)

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
