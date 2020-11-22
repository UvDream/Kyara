package blog

import (
	"github.com/gin-gonic/gin"
	"server/global/response"
	"server/model/request"
	"server/service/blog"
)

func RandomAvatar(c *gin.Context)  {
	var R request.AvatarResponse
	if c.Query("format")=="" {
		R.Format="json"
	}else{
		R.Format= c.Query("format")
	}
	if c.Query("sort")=="" {
		R.Sort="男"
	}else {
		R.Sort=c.Query("sort")
	}
	err,img,msg:=blog.RandomAvatar(R)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(img, msg, c)
	}

}