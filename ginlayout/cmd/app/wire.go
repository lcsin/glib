//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lcsin/glib/ginlayout/internal/handler"
	"github.com/lcsin/glib/ginlayout/internal/repository"
	"github.com/lcsin/glib/ginlayout/internal/repository/dao"
	"github.com/lcsin/glib/ginlayout/internal/service"
	"github.com/lcsin/glib/ginlayout/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		thirdProvider, bookHandlerProvider,
	)

	return gin.Default()
}

// 第三方依赖注入
var thirdProvider = wire.NewSet(
	ioc.InitDB, ioc.InitDBGen, ioc.InitRedis, ioc.InitRouter,
	handler.InitHandlers, handler.InitMiddlewares,
)

// bookHandlerProvider 依赖注入
var bookHandlerProvider = wire.NewSet(
	dao.NewBookDAO,
	//cache.NewUserCache,
	repository.NewBookRepository,
	service.NewBookService,
	handler.NewBookHandler,
)
