package v1

import (
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	service "gin-vue-admin/service/images"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetImagesToken 获取白熊图床token
func GetImagesToken(c *gin.Context) {
	var R request.ImagesStruct
	_ = c.ShouldBindJSON(&R)
	msg := service.GetImagesToken(R)
	response.OkDetailed(msg,"获取成功",c)
}
// 获取图片列表
func GetImagesList(c *gin.Context)  {
	var R request.ImagesListStruct
	_ = c.ShouldBindJSON(&R)
	msg,err:=service.GetImagesList(R)
	if err!=nil {
		response.FailWithMessage("获取错误",c)
	}
	c.JSON(http.StatusOK,msg)
}
