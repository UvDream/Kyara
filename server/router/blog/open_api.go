package blog

import (
	"github.com/gin-gonic/gin"
	"server/api/v1/blog"
)

func InitOpenApi(Router *gin.RouterGroup)(R gin.IRouter){
	OpenApi:=Router.Group("openApi")
	{
		OpenApi.GET("avatar",blog.RandomAvatar)//随机头像
	}
	return OpenApi
}