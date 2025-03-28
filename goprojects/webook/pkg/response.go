package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response[T any] struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ResponseOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response[any]{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
}

func ResponseError(c *gin.Context, code int64, message string) {
	c.JSON(http.StatusOK, Response[any]{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
