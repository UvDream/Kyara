package file

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/model/common/response"
)

type ImageApi struct {
}

// Upload 图片上传
func (i *ImageApi) Upload(c *gin.Context) {
	//TODO
	_, _, err := c.Request.FormFile("file")
	if err != nil {
		response.FailResponse(code.ErrorImageNotFound, c)
		return
	}
	imageService.UploadImageService(c * gin.Context)
}

// Delete 图片删除
func (i *ImageApi) Delete(c *gin.Context) {
	//TODO
}

// List 图片列表
func (i *ImageApi) List(c *gin.Context) {
	//TODO
}
