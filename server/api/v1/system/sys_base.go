package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model/common/response"
	"server/model/system/request"
)

type BaseApi struct{}

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
	if err := b.VerifyCaptcha(loginRequest.Captcha, loginRequest.CaptchaId); err != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "验证码错误",
		})
		return
	}
	//最后验证用户名和密码是否正确
	if msg, err := userService.Login(loginRequest.Username, loginRequest.Password); err != nil {
		response.FailWithMessage(msg, c)
		return
	} else {
		response.OkWithMessage(msg, c)
	}
}
func (b *BaseApi) Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}
