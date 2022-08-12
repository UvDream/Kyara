package article

import (
	"github.com/gin-gonic/gin"
	"server/models/article"
	"server/models/common/response"
)

type TagsApi struct {
}

//CreateTag 创建tag
//@Summary 创建tag
//@Description 创建tag
//@Tags tag
//@Accept json
//@Produce json
//@Param body body article.Tag true "创建tag"
//@Success 200 {object} response.Response"{"code":200,"data":article.Tag,"msg":"操作成功"}"
//@Router /tag/create [post]
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
//@Summary 删除tag
//@Description 删除tag
//@Tags tag
//@Accept  json
//@Produce  json
//@Param id path string true "删除tag"
//@Success 200 {object} response.Response"{"code":200,"data":article.Tag,"msg":"操作成功"}"
//@Router /tag/delete [delete]
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
//@Summary 修改tag
//@Description 修改tag
//@Tags tag
//@Accept  json
//@Produce  json
//@Param body body article.Tag true "修改tag"
//@Success 200 {object} response.Response"{"code":200,"data":article.Tag,"msg":"操作成功"}"
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
//@Summary 查询tag
//@Description 查询tag
//@Tags tag
//@Accept  json
//@Produce  json
//@Success 200 {object} response.Response"{"code":200,"data":[]article.Tag,"msg":"操作成功"}"
//@Router /tag/list [get]
func (t *TagsApi) GetTags(c *gin.Context) {
	keyword := c.Query("keyword")
	list, msg, err := tagService.GetTagsService(keyword)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithDetailed(list, msg, c)

}
