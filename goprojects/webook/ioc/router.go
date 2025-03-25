package ioc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/internal/handler"
)

func InitRouter(middlewares []gin.HandlerFunc, handlers []handler.IHandler) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares...)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 注册路由
	for _, h := range handlers {
		h.RegisterRoutes(r)
	}

	return r
}
