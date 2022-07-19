package theme

import (
	"github.com/gin-gonic/gin"
	code2 "server/code"
	"server/model/common/response"
	"server/model/theme"
	"server/service"
	"server/utils"
)

type ApiThemeGroup struct {
	themeApi
}

var (
	themeService = service.ServicesGroupApp.ThemeServiceGroup
)

type themeApi struct{}

// List 主题列表
func (*themeApi) List(c *gin.Context) {
	keyword := c.Query("keyword")
	xToken := c.Request.Header.Get("x-token")
	j := utils.NewJWT()
	claims, err := j.ParseToken(xToken)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, code, err := themeService.GetThemeListService(keyword, claims.UUID)
	if err != nil {
		response.FailResponse(code, c)
		return
	}
	response.SuccessResponse(data, code, c)
}

// Create 主题创建
func (*themeApi) Create(c *gin.Context) {
	var themeRequest theme.Theme
	err := c.ShouldBindJSON(&themeRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	xToken := c.Request.Header.Get("x-token")
	j := utils.NewJWT()
	claims, err := j.ParseToken(xToken)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	themeRequest.AuthID = claims.UUID
	data, code, err := themeService.CreateThemeService(themeRequest)
	if err != nil {
		response.FailResponse(code, c)
		return
	}
	response.SuccessResponse(data, code, c)
}

// Delete 主题删除
func (*themeApi) Delete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailResponse(code2.ErrorMissingId, c)
		return
	}
	data, code, err := themeService.DeleteThemeService(id)
	if err != nil {
		response.FailResponse(code, c)
		return
	}
	response.SuccessResponse(data, code, c)
}

//Update 主题更新
func (*themeApi) Update(c *gin.Context) {
	var themeRequest theme.Theme
	err := c.ShouldBindJSON(&themeRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if themeRequest.ID == 0 {
		response.FailResponse(code2.ErrorMissingId, c)
		return
	}
	data, code, err := themeService.UpdateThemeService(themeRequest)
	if err != nil {
		response.FailResponse(code, c)
		return
	}
	response.SuccessResponse(data, code, c)
}

//PublicList 查询公开主题
func (*themeApi) PublicList(c *gin.Context) {
	keyword := c.Query("keyword")
	xToken := c.Request.Header.Get("x-token")
	j := utils.NewJWT()
	claims, err := j.ParseToken(xToken)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, code, err := themeService.GetPublicThemeListService(keyword, claims.UUID)
	if err != nil {
		response.FailResponse(code, c)
		return
	}
	response.SuccessResponse(data, code, c)
}

// AdminList 获取所有主题列表
func (*themeApi) AdminList(c *gin.Context) {

}
