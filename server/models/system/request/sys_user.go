package request

import "server/models"

//LoginRequest 登陆请求参数
type LoginRequest struct {
	Username  string `json:"user_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
	CaptchaId string `json:"captcha_id" binding:"required"`
}

// SysUserRequest 用户列表请求参数
type SysUserRequest struct {
	models.PaginationRequest
	Username string `form:"username" json:"user_name"`
	Nickname string `form:"nickname" json:"nick_name"`
}
