package system

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model/system"
	"server/utils"
)

// InitData 初始化数据
// @Summary 初始化数据
// @Description 初始化数据
// @Tags system
// @ID InitData
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response"{"code":200,"data":{},"msg":"ok"}"
// @Router /base/init_data [get]
func (b *BaseApi) InitData(c *gin.Context) {
	initializeUser := []system.User{
		{
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
