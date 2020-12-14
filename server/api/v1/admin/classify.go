package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	resp "server/global/response"
	"server/model/response"
	service "server/service/admin"
	"server/utils"
)

//编辑或者删除分类
func EditClassify(c *gin.Context) {
	var R response.EditClassifyResponse
	err := c.ShouldBindJSON(&R)

	ApiVerify := utils.Rules{
		"TypeName": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(R, ApiVerify)
	if ApiVerifyErr != nil {
		resp.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}

	err, msg, data := service.EditClassify(R)
	if err != nil {
		resp.FailWithMessage(msg, c)
	} else {
		resp.OkDetailed(data, msg, c)
	}
}

//删除分类
func DeleteClassify(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	if id == "" {
		resp.FailWithMessage("缺少分类id", c)
	}
	err, msg := service.DeleteClassify(c)
	if err != nil {
		resp.FailWithMessage(msg, c)
	} else {
		resp.OkWithMessage(msg, c)
	}
}

//获取阿里分类图标
func GetIconfontClassify(c *gin.Context) {
	service.GetIconfontClassify()
}
