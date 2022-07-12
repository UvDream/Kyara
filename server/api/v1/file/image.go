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
	_, _, err := c.Request.FormFile("file")
	if err != nil {
		response.FailResponse(code.ErrorImageNotFound, c)
		return
	}
	data, ce, err := imageService.UploadImageService(c)
	if err != nil {
		response.FailResponse(ce, c)
		return
	}
	response.SuccessResponse(data, ce, c)
}

// Delete 图片删除
func (i *ImageApi) Delete(c *gin.Context) {
	//TODO
}

// List 图片列表
func (i *ImageApi) List(c *gin.Context) {
	//TODO
}
