package file

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type ImageRouter struct{}

func (i *ImageRouter) InitImageRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	imageRouter := Router.Group("image")
	imageApi := v1.ApiGroupApp.FileApiGroup.ImageApi
	{
		imageRouter.POST("/upload", imageApi.Upload)   //图片上传
		imageRouter.DELETE("/delete", imageApi.Delete) //图片删除
		imageRouter.GET("/list", imageApi.List)        //图片列表
	}
	return imageRouter
}
