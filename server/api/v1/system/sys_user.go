package system

import (
	"github.com/gin-gonic/gin"
	"server/models"
	requestModel "server/models/system/request"
)

type UserApi struct{}

func (b *UserApi) UserList(c *gin.Context) {
	var userRequest requestModel.SysUserRequest
	//先判断参数是否合法
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		models.FailWithMessage(err.Error(), c)
	}
	userList, total, msg, err := userService.GetUserListService(&userRequest)
	if err != nil {
		models.FailWithMessage(msg, c)
	}
	models.OkWithDetailed(gin.H{
		"list":  userList,
		"total": total,
	}, msg, c)
}
