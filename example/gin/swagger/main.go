package main

import (
	"net/http"

	_ "gin-swagger/docs"
	"gin-swagger/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	handler.RegisterSwagger(r)

	v1 := r.Group("/api/v1")
	handler.RegisterUserHandler(v1)

	r.Run(":8080")
}
