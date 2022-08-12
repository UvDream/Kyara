package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/code"
	"server/model/common/response"
	"server/model/system"
	"server/model/system/request"
)

type BaseApi struct{}

// Login 登录
// @Tags User
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.LoginRequest true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data=system.User,code=int,msg=string}
// @Router  /public/base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var loginRequest request.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	//首先验证参数是否齐全
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	//其次验证验证码是否正确
	//if err := b.VerifyCaptcha(loginRequest.Captcha, loginRequest.CaptchaId); err != true {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 400,
	//		"msg":  "验证码错误",
	//	})
	//	return
	//}
	//最后验证用户名和密码是否正确
	if user, token, msg, err := userService.Login(loginRequest.Username, loginRequest.Password); err != nil {
		response.FailWithMessage(msg, c)
		return
	} else {
		response.OkWithDetailed(gin.H{"token": token, "user_info": user}, msg, c)
	}
}

// Register 注册
// @Tags User
// @Summary 用户注册
// @Produce  application/json
// @Param data body system.User true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data=system.User,code=int,msg=string}
// @Router  /public/base/register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var registerRequest system.User
	err := c.ShouldBindJSON(&registerRequest)
	if err != nil {
		response.FailResponse(code.ErrorRegisterMissingParam, c)
		return
	}
	data, cd, err := userService.RegisterService(registerRequest)
	if err != nil {
		response.FailResponse(cd, c)
		return
	}
	response.SuccessResponse(data, cd, c)
}
