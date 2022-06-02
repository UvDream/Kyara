package system

import "github.com/gin-gonic/gin"

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
	})
}
