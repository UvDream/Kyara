package blog

import (
	"github.com/gin-gonic/gin"
	"server/global/response"
	"server/model"
	"server/service/blog"
	"server/utils"
)

//博客留言
func Comment(c *gin.Context) {
	var R model.BlogComment
	_ = c.ShouldBindJSON(&R)
	ApiVerify := utils.Rules{
		"CommentContent": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(R, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	blog.Comment(R)
}
