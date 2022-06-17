package article

import (
	"github.com/gin-gonic/gin"
	"server/model/article"
	"server/model/common/response"
)

type TagsApi struct {
}

//CreateTag 创建tag
func (t *TagsApi) CreateTag(c *gin.Context) {
	var tag article.Tag
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	list, msg, err := tagService.CreateTagService(tag)
	if err != nil {
		response.FailWithMessage(msg, c)
	}
	response.OkWithDetailed(list, msg, c)
}

//DeleteTag 删除tag
func (t *TagsApi) DeleteTag(c *gin.Context) {

}

//UpdateTag 修改tag
func (t *TagsApi) UpdateTag(c *gin.Context) {

}

//GetTags 查询tag
func (t *TagsApi) GetTags(c *gin.Context) {

}
