package system

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("register", baseApi.Register)
	}
	return userRouter
}
