package theme

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type ThemesGroupRouter struct {
}

func (*ThemesGroupRouter) InitThemeRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	themeRouter := Router.Group("theme")
	themeApi := v1.ApiGroupApp.ThemeApiGroup
	{
		themeRouter.GET("/list", themeApi.List)
		themeRouter.GET("/public", themeApi.PublicList)
		themeRouter.GET("/detail", themeApi.Detail)
		themeRouter.POST("/create", themeApi.Create)
		themeRouter.DELETE("/delete", themeApi.Delete)
		themeRouter.PUT("/update", themeApi.Update)
		themeRouter.GET("/admin/list", themeApi.AdminList)
	}
	return themeRouter
}
