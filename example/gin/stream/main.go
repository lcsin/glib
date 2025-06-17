package main

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 批量上传接口
	r.POST("/batch-upload", func(c *gin.Context) {
		// 创建SSE流
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		form, _ := c.MultipartForm()
		files := form.File["files"]
		total := len(files)

		for i, file := range files {
			// 处理单个文件
			processFile(file)

			// 发送进度事件
			c.SSEvent("progress", gin.H{
				"current": i + 1,
				"total":   total,
			})
			c.Writer.Flush()
		}

		c.SSEvent("complete", nil)
	})

	r.Run(":8080")
}

func processFile(file *multipart.FileHeader) {
	// 实际文件处理逻辑
	fmt.Println("success !!!")
}
