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
	if R.ID==0{
		ApiVerify := utils.Rules{
			"UserName":{utils.NotEmpty()},
			"Email":{utils.NotEmpty()},
			"Avatar":{utils.NotEmpty()},
			"CommentContent": {utils.NotEmpty()},
		}
		ApiVerifyErr := utils.Verify(R, ApiVerify)
		if ApiVerifyErr != nil {
			response.FailWithMessage(ApiVerifyErr.Error(), c)
			return
		}
	}

	err,msg:=blog.Comment(R)
	if err!=nil{
		response.FailWithMessage(msg,c)
	}else {
		response.OkWithMessage(msg,c)
	}
}
