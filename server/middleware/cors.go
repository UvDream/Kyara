package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/config"
	"server/global"
)

// Cors 跨域配置(放行所有的方法)
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func CorsByRules() gin.HandlerFunc {
	if global.Config.Cors.Mode == "allow-all" {
		return Cors()
	}
	return func(c *gin.Context) {
		whiteList := checkCors(c.GetHeader("origin"))
		if whiteList != nil {
			c.Header("Access-Control-Allow-Origin", whiteList.AllowOrigin)
			c.Header("Access-Control-Allow-Headers", whiteList.AllowHeaders)
			c.Header("Access-Control-Allow-Methods", whiteList.AllowMethods)
			c.Header("Access-Control-Expose-Headers", whiteList.ExposeHeaders)
			if whiteList.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}
		// 严格白名单模式且未通过检查，直接拒绝处理请求
		if whiteList == nil && global.Config.Cors.Mode == "strict-whitelist" && !(c.Request.Method == "GET" && c.Request.URL.Path == "/health") {
			c.AbortWithStatus(http.StatusForbidden)
		} else {
			// 非严格白名单模式，无论是否通过检查均放行所有 OPTIONS 方法
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
		}
		c.Next()

	}
}
func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range global.Config.Cors.Whitelist {
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
