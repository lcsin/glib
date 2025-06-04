package ilogrus

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func NewErrorHook() (*ErrorHook, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fp := path.Join(dir, "err")
	if err = os.MkdirAll(fp, os.ModePerm); err != nil {
		return nil, err
	}

	//fp := fmt.Sprintf("err_%v.log", time.Now().Format("20060102"))
	//file, err := os.OpenFile(path.Join(dir, fp), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	return nil, err
	//}

	return &ErrorHook{dir: fp}, nil
}

type ErrorHook struct {
	writer *os.File
	dir    string
	sync.Once
}

// Levels 声明钩子适用的日志级别
func (h *ErrorHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
	}
}

// Fire 在日志记录时执行，添加函数名信息
func (h *ErrorHook) Fire(entry *logrus.Entry) error {
	fp := fmt.Sprintf("err_%v.log", time.Now().Format("20060102"))
	f, err := os.OpenFile(path.Join(h.dir, fp), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	h.writer = f

	_, file, line, ok := runtime.Caller(8)
	if ok {
		entry.Data["caller"] = file
		entry.Data["line"] = line
	}

	msg, err := entry.String()
	if err != nil {
		return err
	}
	_, err = h.writer.WriteString(msg)

	return err
}

func getCallers() []string {
	var callers []string
	for i := 2; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		callers = append(callers, fmt.Sprintf("%v/%v", file, line))
	}
	return callers
}
