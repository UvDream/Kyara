package admin

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

// 图床相关路由
func InitOuterImages(Router *gin.RouterGroup) (R gin.IRoutes) {
	//ImagesRouter:=Router.Group("images").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	ImagesRouter := Router.Group("images")
	{
		ImagesRouter.POST("getToken", v1.GetImagesToken)     // 白熊图床获取token
		ImagesRouter.POST("getImagesList", v1.GetImagesList) //图床获取图片列表
		ImagesRouter.POST("uploadImage", v1.UploadImage)     //上传图片
	}
	return ImagesRouter
}
