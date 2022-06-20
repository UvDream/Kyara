package article

import (
	"github.com/gin-gonic/gin"
	"server/model/article"
	"server/model/common/response"
)

type CategoriesApi struct{}

//CreateCategory 创建Category
func (ct *CategoriesApi) CreateCategory(c *gin.Context) {
	var category article.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, msg, err := categoryService.CreateCategoryService(category); err != nil {
		response.FailWithMessage(msg, c)
		return
	} else {
		response.OkWithDetailed(list, msg, c)
	}
}

//DeleteCategory 删除Category
func (ct *CategoriesApi) DeleteCategory(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailWithMessage("id不能为空", c)
		return
	}
	if msg, err := categoryService.DeleteCategoryService(id); err != nil {
		response.FailWithMessage(msg, c)
		return
	} else {
		response.OkWithMessage(msg, c)
	}
}

//UpdateCategory 修改Category
func (ct *CategoriesApi) UpdateCategory(c *gin.Context) {
	var category article.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if category.ID == 0 {
		response.FailWithMessage("id不能为空", c)
		return
	}
	if msg, err := categoryService.UpdateCategoryService(category); err != nil {
		response.FailWithMessage(msg, c)
		return
	} else {
		response.OkWithMessage(msg, c)
	}

}

//GetCategories 查询Category
func (ct *CategoriesApi) GetCategories(c *gin.Context) {
	//keyword := c.Query("keyword")
	if list, msg, err := categoryService.GetCategoryService(); err != nil {
		response.FailWithMessage(msg, c)
		return
	} else {
		response.OkWithDetailed(list, msg, c)
	}
}
