package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 注册pprof路由
	pprof.Register(r)

	r.Run()
}
