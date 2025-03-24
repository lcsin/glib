package iutil

import (
	"archive/zip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// PathExists 判断文件或目录是否存在
func PathExists(src string) bool {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return false
	}
	return true
}

// CopyFile 拷贝文件
func CopyFile(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	buf, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, buf, info.Mode())
}

// CopyDir 拷贝目录
func CopyDir(src, dst string, ignores []string) error {
	srcinfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, srcinfo.Mode())
	if err != nil {
		return err
	}

	fds, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	hasIgnores := func(name string, ignores []string) bool {
		for _, v := range ignores {
			if v == name {
				return true
			}
		}
		return false
	}

	for _, fd := range fds {
		if hasIgnores(fd.Name(), ignores) {
			continue
		}
		srcfp := filepath.Join(src, fd.Name())
		dstfp := filepath.Join(dst, fd.Name())
		if fd.IsDir() {
			err = CopyDir(srcfp, dstfp, ignores)
		} else {
			err = CopyFile(srcfp, dstfp)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// FileMD5 获取文件MD5
func FileMD5(fp string) (string, error) {
	file, err := os.Open(fp)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)[:16])), nil
}

// ZipWithCompress 压缩指定文件集
func ZipWithCompress(files []string, des string, abs bool) error {
	dir := filepath.Dir(des)
	if !PathExists(dir) && os.MkdirAll(dir, os.ModePerm) != nil {
		return fmt.Errorf("dir not exist")
	}

	zipFile, err := os.Create(des)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, f := range files {
		if err = zipWithCompress(zipWriter, f, abs); err != nil {
			return err
		}
	}
	return nil
}

// 添加文件到zip
func zipWithCompress(w *zip.Writer, fp string, abs bool) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	if abs {
		header.Name = fp
	}

	writer, err := w.CreateHeader(header)
	if err != nil {
		return err
	}

	if _, err = io.Copy(writer, f); err != nil {
		return err
	}
	return nil
}

// Unzip 解压zip包
func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer reader.Close()

	if err = os.MkdirAll(target, 0755); err != nil {
		return err
	}

	unzip := func(file *zip.File) error {
		fp := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			_ = os.MkdirAll(fp, file.Mode())
			return nil
		}

		dir := filepath.Dir(fp)
		if len(dir) > 0 {
			if _, err = os.Stat(dir); os.IsNotExist(err) {
				err = os.MkdirAll(dir, 0755)
				if err != nil {
					return err
				}
			}
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err = io.Copy(targetFile, fileReader); err != nil {
			return err
		}

		return nil
	}

	for _, file := range reader.File {
		if err = unzip(file); err != nil {
			return err
		}
	}

	return nil
}
