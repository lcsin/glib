package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 同时将日志写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		//fmt.Fprintln(gin.DefaultWriter, "got ping return pong ...")
		c.String(200, "pong")
	})

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
	})

	router.Run()
}
