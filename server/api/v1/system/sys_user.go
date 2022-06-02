package system

import "github.com/gin-gonic/gin"

type UserApi struct{}

func (b *UserApi) Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
