package v1

import (
	"gin-vue-admin/model/request"
	service "gin-vue-admin/service/images"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetImagesToken 获取图床token
func GetImagesToken(c *gin.Context) {
	var R request.ImagesStruct
	_ = c.ShouldBindJSON(&R)
	R.Email = "913906565@qq.com"
	R.Password = "11165wzj"
	msg := service.GetImagesToken(R)
	c.JSON(http.StatusOK, msg)
}
