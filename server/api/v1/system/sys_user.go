package system

import (
	"github.com/gin-gonic/gin"
	"server/models/common/response"
	requestModel "server/models/system/request"
)

type UserApi struct{}

func (b *UserApi) UserList(c *gin.Context) {
	var userRequest requestModel.SysUserRequest
	//先判断参数是否合法
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	userList, total, msg, err := userService.GetUserListService(&userRequest)
	if err != nil {
		response.FailWithMessage(msg, c)
	}
	response.OkWithDetailed(gin.H{
		"list":  userList,
		"total": total,
	}, msg, c)
}
