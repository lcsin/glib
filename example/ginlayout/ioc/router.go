package ioc

import (
	"net/http"

	"ginlayout/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares []gin.HandlerFunc, handlers []handler.IHandler) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares...)
	v1 := r.Group("/ap/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 注册路由
	for _, h := range handlers {
		h.RegisterRoutes(v1)
	}

	return r
}
