package theme

import (
	"github.com/gin-gonic/gin"
	code2 "server/code"
	"server/models/common/response"
	"server/models/theme"
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
// @Summary 主题列表
// @Description 主题列表
// @Tags theme
// @Accept  json
// @Produce  json
// @Param keyword query string false "关键词"
// @Success 200 {object} response.Response "{"code":200,"data":[]theme.ResponseTheme,"msg":"ok"}"
// @Router /theme/list [get]
func (*themeApi) List(c *gin.Context) {
	keyword := c.Query("keyword")
	xToken := c.Request.Header.Get("x-token")
	j := utils.NewJWT()
	claims, err := j.ParseToken(xToken)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, code, err := themeService.GetThemeListService(keyword, claims.ID)
	if err != nil {
		response.FailResponse(code, c)
		return
	}
	response.SuccessResponse(data, code, c)
}

// Create 主题创建
// @Summary 主题创建
// @Description 主题创建
// @Tags theme
// @Accept  json
// @Produce  json
// @Param theme body theme.Theme true "主题"
// @Success 200 {object} response.Response "{"code":200,"data":theme.ResponseTheme,"msg":"ok"}"
// @Router /theme/create [post]
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
	themeRequest.UserID = claims.ID
	data, code, err := themeService.CreateThemeService(themeRequest)
	if err != nil {
		response.FailResponse(code, c)
		return
	}
	response.SuccessResponse(data, code, c)
}

// Delete 主题删除
// @Summary 主题删除
// @Description 主题删除
// @Tags theme
// @Accept  json
// @Produce  json
// @Param id query string true "主题id"
// @Success 200 {object} response.Response "{"code":200,"data":theme.Theme,"msg":"ok"}"
// @Router /theme/delete [delete]
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
// @Summary 主题更新
// @Description 主题更新
// @Tags theme
// @Accept  json
// @Produce  json
// @Param theme body theme.Theme true "主题"
// @Success 200 {object} response.Response "{"code":200,"data":theme.Theme,"msg":"ok"}"
// @Router /theme/update [put]
func (*themeApi) Update(c *gin.Context) {
	var themeRequest theme.Theme
	err := c.ShouldBindJSON(&themeRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if themeRequest.ID == "" {
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
// @Summary 查询公开主题
// @Description 查询公开主题
// @Tags theme
// @Accept  json
// @Produce  json
// @Param keyword query string false "关键词"
// @Success 200 {object} response.Response "{"code":200,"data":[]theme.ResponseTheme,"msg":"ok"}"
// @Router /theme/public [get]
func (*themeApi) PublicList(c *gin.Context) {
	keyword := c.Query("keyword")
	xToken := c.Request.Header.Get("x-token")
	j := utils.NewJWT()
	claims, err := j.ParseToken(xToken)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, code, err := themeService.GetPublicThemeListService(keyword, claims.ID)
	if err != nil {
		response.FailResponse(code, c)
		return
	}
	response.SuccessResponse(data, code, c)
}

// Detail 主题详情
// @Summary 主题详情
// @Description 主题详情
// @Tags theme
// @Accept  json
// @Produce  json
// @Param id path string true "主题id"
// @Success 200 {object} response.Response "{"code":200,"data":theme.Theme,"msg":"ok"}"
// @Router /theme/detail [get]
func (*themeApi) Detail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailResponse(code2.ErrorMissingId, c)
		return
	}
	data, code, err := themeService.GetThemeDetailService(id)
	if err != nil {
		response.FailResponse(code, c)
		return
	}
	response.SuccessResponse(data, code, c)
}

// AdminList 获取所有主题列表
func (*themeApi) AdminList(c *gin.Context) {

}
