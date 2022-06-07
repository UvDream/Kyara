package system

import "github.com/gin-gonic/gin"

type UserApi struct{}

func (b *UserApi) UserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
