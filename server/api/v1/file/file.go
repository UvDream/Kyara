package file

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/model/common/request"
	"server/model/common/response"
	"server/utils"
	"strconv"
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
	var fileOpts request.PaginationRequest
	fileOpts.KeyWord = c.Query("key_word")
	fileOpts.Page, _ = strconv.Atoi(c.Query("page"))
	fileOpts.PageSize, _ = strconv.Atoi(c.Query("page_size"))
	j := utils.NewJWT()
	xToken := c.Request.Header.Get("x-token")
	claims, err := j.ParseToken(xToken)
	if err != nil {
		response.FailWithMessage("token验证失败", c)
		return
	}
	list, total, ce, err := fileService.ListFileService(fileOpts, claims.UUID)
	if err != nil {
		response.FailResponse(ce, c)
		return
	}
	response.SuccessResponse(map[string]interface{}{
		"list":  list,
		"total": total,
	}, ce, c)
}
