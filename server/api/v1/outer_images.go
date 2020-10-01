package v1

import (
	"fmt"
	"server/global/response"
	"server/model/request"
	service "server/service/images"
	"net/http"
	"github.com/gin-gonic/gin"
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
	_, header, err := c.Request.FormFile("image")
	fmt.Println(header,err)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
	}
	_,err=service.UploadImage(c)
	if err!=nil{
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
	}else{
		response.OkWithMessage("上传成功",c)
	}
}