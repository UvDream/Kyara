package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/global/response"
	"server/model/request"
	service "server/service/images"
)

// GetImagesToken 获取白熊图床token
func GetImagesToken(c *gin.Context) {
	var R request.ImagesStruct
	_ = c.ShouldBindJSON(&R)
	 msg ,err:=service.GetImagesToken(R)
	if err!=nil {
		response.FailWithMessage(msg,c)
	}else{
		response.OkWithMessage(msg,c)
	}

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
//上传图片
func UploadImage(c *gin.Context)  {
	_, _, err := c.Request.FormFile("image")
	if err != nil {
		response.FailWithMessage("缺少上传图片", c)
	}
	msg,err:=service.UploadImage(c)
	if err!=nil{
		response.FailWithMessage(msg, c)
	}else{
		response.OkWithMessage(msg,c)
	}
} 