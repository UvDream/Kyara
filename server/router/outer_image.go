package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)
// 图床相关路由
func InitOuterImages(Router *gin.RouterGroup) (R gin.IRoutes){
	//ImagesRouter:=Router.Group("images").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	ImagesRouter:=Router.Group("images")
	{
		ImagesRouter.POST("getToken",v1.GetImagesToken) // 白熊图床获取token
		ImagesRouter.POST("getImagesList",v1.GetImagesList)
	}
	return ImagesRouter
}