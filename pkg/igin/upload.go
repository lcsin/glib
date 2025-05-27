package api

import (
	"fmt"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/glib/pkg/iutil"
)

// UploadSingleFile 单文件上传
func UploadSingleFile(c *gin.Context, basePath string) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	filename := RenameUploadFile(file.Filename)
	fp := path.Join(basePath, filename)
	if err = c.SaveUploadedFile(file, fp); err != nil {
		return err
	}
	return nil
}

// UploadMultiFile 多文件上传
func UploadMultiFile(c *gin.Context, basePath string) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["files"]
	for _, f := range files {
		filename := RenameUploadFile(f.Filename)
		fp := path.Join(basePath, filename)
		if err = c.SaveUploadedFile(f, fp); err != nil {
			return err
		}
	}

	return nil
}

// RenameUploadFile 重命名上传文件
func RenameUploadFile(origin string) string {
	return fmt.Sprintf("%v_%v%v",
		iutil.RandomString(5),
		time.Now().Format("20060102150405"),
		path.Ext(origin))
}
