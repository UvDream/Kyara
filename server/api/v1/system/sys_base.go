package system

import "github.com/gin-gonic/gin"

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
