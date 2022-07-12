package file

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type FilesRouter struct{}

func (i *FilesRouter) InitFileRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	fileRouter := Router.Group("file")
	fileApi := v1.ApiGroupApp.FileApiGroup.FilesApi
	{
		fileRouter.POST("/upload", fileApi.Upload) //图片上传
	}
	return fileRouter
}
