package article

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type CategoriesStruct struct {
}

func (t *CategoriesStruct) InitCategoriesRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	categoriesRouter := Router.Group("category")
	categoriesApi := v1.ApiGroupApp.CategoryApiGroup
	{
		categoriesRouter.POST("/create", categoriesApi.CreateCategory)
		categoriesRouter.DELETE("/delete", categoriesApi.DeleteCategory)
		categoriesRouter.PUT("/update", categoriesApi.UpdateCategory)
		categoriesRouter.GET("/list", categoriesApi.GetCategories)
	}
	return categoriesRouter
}
