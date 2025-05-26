package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	RegisterRoutes(v1 *gin.RouterGroup)
}

func InitHandlers(helloHandler *HelloHandler) []IHandler {
	return []IHandler{
		helloHandler,
	}
}

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		// 跨域中间件
		cors.Default(),
	}
}
