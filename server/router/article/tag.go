package article

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type TagsStruct struct {
}

func (t *TagsStruct) InitTagRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	tagRouter := Router.Group("tag")
	tagApi := v1.ApiGroupApp.TagApiGroup
	{
		tagRouter.POST("/create", tagApi.CreateTag)
		tagRouter.DELETE("/delete", tagApi.DeleteTag)
		tagRouter.PUT("/update", tagApi.UpdateTag)
		tagRouter.GET("/list", tagApi.GetTags)
	}
	return tagRouter
}
