package v1

import (
	"github.com/gin-gonic/gin"
	"server/global/response"
	"server/model"
	service "server/service/admin"
	"server/utils"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var R model.SysArticle
	err := c.ShouldBindJSON(&R)
	ApiVerify := utils.Rules{
		"Title": {utils.NotEmpty()},
		"ArticleContent": {utils.NotEmpty()},
		"ClassificationID": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(R, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err, msg, data := service.AddArticle(R, c)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(data, msg, c)
	}
}

//查询文章详情
func GetArticleDetail(c *gin.Context) {
	id := c.Query("id")
	err, data := service.GetArticleDetail(id)
	if err != nil {
		response.FailWithMessage("获取详情错误", c)
	} else {
		response.OkDetailed(data, "获取详情成功", c)
	}
}

//增加公告
func AddNotice(c *gin.Context) {
	var R model.BlogNotice
	err := c.ShouldBindJSON(&R)
	ApiVerify := utils.Rules{
		"Title":   {utils.NotEmpty()},
		"Content": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(R, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err, msg := service.AddNotice(R)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkWithMessage(msg, c)
	}
}

//增加tag
func AddTag(c *gin.Context) {
	if c.Query("tag") == "" {
		response.FailWithMessage("缺少tag", c)
	}
	err, list, msg := service.AddTag(c)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(list, msg, c)
	}
}
