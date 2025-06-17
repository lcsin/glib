package api

import (
	"fmt"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/glib/pkg/iutil"
)

type FileInfo struct {
	Name   string
	Path   string
	Size   int64
	Suffix string
}

// UploadSingleFile 单文件上传
func UploadSingleFile(c *gin.Context, basePath string) (*FileInfo, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	filename := RenameUploadFile(file.Filename)
	fp := path.Join(basePath, filename)
	if err = c.SaveUploadedFile(file, fp); err != nil {
		return nil, err
	}

	return &FileInfo{
		Name:   file.Filename,
		Path:   fp,
		Size:   file.Size,
		Suffix: path.Ext(file.Filename),
	}, nil
}

// UploadMultiFile 多文件上传
func UploadMultiFile(c *gin.Context, basePath string) ([]*FileInfo, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := form.File["files"]
	fileList := make([]*FileInfo, 0, len(files))
	for _, f := range files {
		filename := RenameUploadFile(f.Filename)
		fp := path.Join(basePath, filename)
		if err = c.SaveUploadedFile(f, fp); err != nil {
			return fileList, err
		}

		fileList = append(fileList, &FileInfo{
			Name:   f.Filename,
			Path:   fp,
			Size:   f.Size,
			Suffix: path.Ext(f.Filename),
		})
	}

	return fileList, nil
}

// RenameUploadFile 重命名上传文件
func RenameUploadFile(origin string) string {
	return fmt.Sprintf("%v_%v%v",
		iutil.RandomString(5),
		time.Now().Format("20060102150405"),
		path.Ext(origin))
}
