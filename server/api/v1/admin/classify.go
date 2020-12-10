package v1

import (
	"github.com/gin-gonic/gin"
	"server/model/response"
	resp "server/global/response"
	service "server/service/admin"
	"server/utils"
)
//编辑或者删除分类
func EditClassify(c *gin.Context){
	var R response.EditClassifyResponse
	err:=c.ShouldBindJSON(&R)

	ApiVerify := utils.Rules{
		"TypeName": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(R, ApiVerify)
	if ApiVerifyErr != nil {
		resp.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}

	err,msg,data:=service.EditClassify(R)
	if err != nil {
		resp.FailWithMessage(msg, c)
	} else {
		resp.OkDetailed(data,msg, c)
	}
}
