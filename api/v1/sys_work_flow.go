package v1

import (
	"fmt"
	"blog-api/global/response"
	"blog-api/model"
	"blog-api/service"
	"blog-api/utils"
	"github.com/gin-gonic/gin"
)

// @Tags workflow
// @Summary 注册工作流
// @Produce  application/json
// @Param data body model.SysWorkflow true "注册工作流接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /workflow/createWorkFlow [post]
func CreateWorkFlow(c *gin.Context) {
	var wk model.SysWorkflow
	_ = c.ShouldBindJSON(&wk)
	WKVerify := utils.Rules{
		"WorkflowNickName":    {utils.NotEmpty()},
		"WorkflowName":        {utils.NotEmpty()},
		"WorkflowDescription": {utils.NotEmpty()},
		"WorkflowStepInfo":    {utils.NotEmpty()},
	}
	WKVerifyErr := utils.Verify(wk, WKVerify)
	if WKVerifyErr != nil {
		response.FailWithMessage(WKVerifyErr.Error(), c)
		return
	}
	err := service.Create(wk)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithMessage("获取成功", c)
	}
}
