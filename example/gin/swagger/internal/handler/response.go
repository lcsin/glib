package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int64  `json:"code" example:"0"`
	Message string `json:"message" example:"ok"`
	Data    any    `json:"data"`
}

func ResponseOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
}

func ResponseError(c *gin.Context, code int64, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
