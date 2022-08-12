package file

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models"
	"server/utils"
	"strconv"
)

type FilesApi struct{}

// Upload 文件上传
// @Tags file
// @Summary 文件上传
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {string} string "{""code"":200,""message"":""上传成功""}"
// @Failure 400 {string} string "{""code"":400,""message"":""上传失败""}"
// @Failure 500 {string} string "{""code"":500,""message"":""服务器错误""}"
// @Router /file/upload [post]
func (i *FilesApi) Upload(c *gin.Context) {
	_, _, err := c.Request.FormFile("file")
	if err != nil {
		models.FailResponse(code.ErrorImageNotFound, c)
		return
	}
	data, ce, err := fileService.UploadFileService(c)
	if err != nil {
		models.FailResponse(ce, c)
		return
	}
	models.SuccessResponse(data, ce, c)
}

// Delete 文件删除
// @Tags file
// @Summary 文件删除
// @Produce json
// @Param    id   query     string  true  "参数"
// @Success 200 {string} string "{""code"":200,""message"":""删除成功""}"
// @Failure 400 {string} string "{""code"":400,""message"":""删除失败""}"
// @Failure 500 {string} string "{""code"":500,""message"":""服务器错误""}"
// @Router /file/delete/{id} [delete]
func (i *FilesApi) Delete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		models.FailResponse(code.ErrorMissingId, c)
		return
	}
	data, ce, err := fileService.DeleteFileService(id)
	if err != nil {
		models.FailResponse(ce, c)
		return
	}
	models.SuccessResponse(data, ce, c)
}

// List 文件列表
// @Tags file
// @Summary 文件列表
// @Produce json
// @Param    query   query     models.PaginationRequest  false  "分页参数"
// @Success 200 {string} string "{""code"":200,""message"":""删除成功""}"
// @Failure 400 {string} string "{""code"":400,""message"":""删除失败""}"
// @Failure 500 {string} string "{""code"":500,""message"":""服务器错误""}"
// @Router /file/list [get]
func (i *FilesApi) List(c *gin.Context) {
	var fileOpts models.PaginationRequest
	fileOpts.KeyWord = c.Query("key_word")
	fileOpts.Page, _ = strconv.Atoi(c.Query("page"))
	fileOpts.PageSize, _ = strconv.Atoi(c.Query("page_size"))
	j := utils.NewJWT()
	xToken := c.Request.Header.Get("x-token")
	claims, err := j.ParseToken(xToken)
	if err != nil {
		models.FailWithMessage("token验证失败", c)
		return
	}
	list, total, ce, err := fileService.ListFileService(fileOpts, claims.ID)
	if err != nil {
		models.FailResponse(ce, c)
		return
	}
	models.SuccessResponse(map[string]interface{}{
		"list":  list,
		"total": total,
	}, ce, c)
}
