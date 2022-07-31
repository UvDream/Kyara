package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/docs"
)

func InitSwag(Router *gin.Engine) {
	docs.SwaggerInfo.Title = "后台API文档示例"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
