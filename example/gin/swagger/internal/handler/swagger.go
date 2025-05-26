package handler

import (
	"gin-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(r gin.IRouter) {
	docs.SwaggerInfo.Title = "gin-swagger"
	docs.SwaggerInfo.Description = "这是一个gin整合swagger的示例"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Version = "v1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
