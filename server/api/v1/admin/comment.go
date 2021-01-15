package v1

import (
	"github.com/gin-gonic/gin"
	resp "server/global/response"
	"server/model/request"
	"server/model/response"
	service "server/service/admin"
)

// 获取留言
func GetBlogComment(c *gin.Context) {
	var R request.ListStruct
	err := c.ShouldBindJSON(&R)
	if err != nil {
		resp.FailWithMessage("获取参数错误", c)
	}
	err, list, msg, total := service.GetBlogComment(R)
	if err != nil {
		resp.FailWithMessage(msg, c)
	} else {
		resp.OkDetailed(&response.PageResult{List: list, Total: total, Page: R.Page, PageSize: R.PageSize}, msg, c)
	}
}

// 审核留言
func CheckBlogComment(c *gin.Context) {
	var R response.CheckComment
	err := c.ShouldBindJSON(&R)
	if err != nil {
		resp.FailWithMessage("获取参数错误", c)
	}
	err, msg := service.CheckBlogComment(R)
	if err != nil {
		resp.FailWithMessage(msg, c)
	} else {
		resp.OkWithMessage(msg, c)
	}
}

//回复留言
func RevertComment(c *gin.Context) {
	var R response.ReplyComment
	err := c.ShouldBindJSON(&R)
	if err != nil {
		resp.FailWithMessage("获取参数错误", c)
	}
	msg, err := service.RevertComment(c, R)
	if err != nil {
		resp.FailWithMessage(msg, c)
	} else {
		resp.OkWithMessage(msg, c)
	}
}

// 删除留言
func DeleteComment(c *gin.Context) {
	msg, err := service.DeleteComment(c)
	if err != nil {
		resp.FailWithMessage(msg, c)
	} else {
		resp.OkWithMessage(msg, c)
	}
}
