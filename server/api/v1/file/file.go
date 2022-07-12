package file

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/model/common/response"
)

type FilesApi struct{}

// Upload 文件上传
func (i *FilesApi) Upload(c *gin.Context) {
	_, _, err := c.Request.FormFile("file")
	if err != nil {
		response.FailResponse(code.ErrorImageNotFound, c)
		return
	}
	data, ce, err := fileService.UploadFileService(c)
	if err != nil {
		response.FailResponse(ce, c)
		return
	}
	response.SuccessResponse(data, ce, c)
}

// Delete 文件删除
func (i *FilesApi) Delete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailResponse(code.ErrorMissingId, c)
		return
	}
	data, ce, err := fileService.DeleteFileService(id)
	if err != nil {
		response.FailResponse(ce, c)
		return
	}
	response.SuccessResponse(data, ce, c)
}

// List 文件列表
func (i *FilesApi) List(c *gin.Context) {
	//TODO
}
