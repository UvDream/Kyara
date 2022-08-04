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
		return
	}
	response.OkWithDetailed(list, msg, c)
}

//DeleteTag 删除tag
func (t *TagsApi) DeleteTag(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailWithMessage("id必填", c)
	}
	msg, err := tagService.DeleteTagService(id)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithMessage(msg, c)
}

//UpdateTag 修改tag
func (t *TagsApi) UpdateTag(c *gin.Context) {
	var tag article.Tag
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if tag.ID == "" {
		response.FailWithMessage("缺少tag id", c)
		return
	}
	data, msg, err := tagService.UpdateTagService(tag)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithDetailed(data, msg, c)
}

//GetTags 查询tag
func (t *TagsApi) GetTags(c *gin.Context) {
	keyword := c.Query("keyword")
	list, msg, err := tagService.GetTagsService(keyword)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithDetailed(list, msg, c)

}
