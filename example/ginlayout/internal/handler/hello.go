package handler

import (
	"net/http"

	"ginlayout/internal/service"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	srv service.IHelloService
}

func NewHelloHandler(srv service.IHelloService) *HelloHandler {
	return &HelloHandler{srv: srv}
}

func (h *HelloHandler) RegisterRoutes(v1 *gin.RouterGroup) {
	v1.GET("/sayHello", h.SayHello)
}

func (h *HelloHandler) SayHello(c *gin.Context) {
	hello, err := h.srv.SayHello(c, 1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "error",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    hello,
	})
}
