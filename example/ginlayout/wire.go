//go:build wireinject

package main

import (
	"ginlayout/internal/handler"
	"ginlayout/internal/repository"
	"ginlayout/internal/repository/cache"
	"ginlayout/internal/repository/dao"
	"ginlayout/internal/service"
	"ginlayout/ioc"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 初始化
		ioc.InitDB, ioc.InitRedis, handler.InitMiddlewares, handler.InitHandlers, ioc.InitRouter,
		// hello service
		dao.NewHelloDAO, cache.NewHelloCache, repository.NewHelloRepository, service.NewHelloService, handler.NewHelloHandler,
	)

	return gin.Default()
}
