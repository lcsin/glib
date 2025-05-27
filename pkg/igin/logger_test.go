package api

import (
	"io"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogger(t *testing.T) {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 同时将日志写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		Info("got ping return pong ...")
		Infof("result: %v, id=%d", "got ping return pong ...", 10086)
		c.String(200, "pong")
	})

	router.Run()
}
