package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
	getCode "server/code"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
}

const (
	ERROR   = 7
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context, success bool) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
		success,
	})
}

// SuccessResponse 成功返回
func SuccessResponse(data interface{}, code int, c *gin.Context) {
	Result(code, data, getCode.Text(code), c, true)
}

// FailResponse 失败返回
func FailResponse(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, getCode.Text(code), c, false)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c, false)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c, false)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c, true)
}
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c, true)
}
