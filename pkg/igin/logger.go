package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Info(a ...any) {
	ginFormat := fmt.Sprintf("[GIN] %v | %s |",
		time.Now().Format("2006/01/02 - 15:04:05"),
		"INFO")
	log := fmt.Sprint(a...)

	fmt.Fprintln(gin.DefaultWriter, ginFormat, log)
}

func Infof(format string, a ...any) {
	ginFormat := fmt.Sprintf("[GIN] %v | %s |",
		time.Now().Format("2006/01/02 - 15:04:05"),
		"INFO")
	log := fmt.Sprintf(format, a...)

	fmt.Fprintln(gin.DefaultWriter, ginFormat, log)
}

func Error(a ...any) {
	ginFormat := fmt.Sprintf("[GIN] %v | %s |",
		time.Now().Format("2006/01/02 - 15:04:05"),
		"ERROR")
	log := fmt.Sprint(a...)

	fmt.Fprintln(gin.DefaultErrorWriter, ginFormat, log)
}

func Errorf(format string, a ...any) {
	ginFormat := fmt.Sprintf("[GIN] %v | %s |",
		time.Now().Format("2006/01/02 - 15:04:05"),
		"ERROR")
	log := fmt.Sprintf(format, a...)

	fmt.Fprintln(gin.DefaultErrorWriter, ginFormat, log)
}
