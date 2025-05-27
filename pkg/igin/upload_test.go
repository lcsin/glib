package api

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestName(t *testing.T) {
	r := gin.Default()

	r.POST("/uploadFile", func(c *gin.Context) {
		if err := UploadSingleFile(c, "file"); err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}
		c.String(http.StatusOK, "ok")
	})

	r.POST("/uploadFiles", func(c *gin.Context) {
		if err := UploadMultiFile(c, "file"); err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}
		c.String(http.StatusOK, "ok")
	})

	r.Run()
}
