package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}
