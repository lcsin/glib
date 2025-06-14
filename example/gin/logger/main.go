package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/glib/pkg/ilogrus"
	"github.com/sirupsen/logrus"
)

func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 同时将日志写入文件和控制台
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//logrus.AddHook(ilogrus.NewErrorHook())

	logrus.AddHook(ilogrus.NewErrorHook())

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		logrus.Error("test error log ...")
		c.String(200, "pong")
	})

	//router.GET("/pong", func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong1")
	//}, func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong2")
	//})

	router.Run()
}
