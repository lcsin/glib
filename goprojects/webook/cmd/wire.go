//go:build wireinject

package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lcsin/webook/internal/handler"
	"github.com/lcsin/webook/internal/repository"
	"github.com/lcsin/webook/internal/repository/cache"
	"github.com/lcsin/webook/internal/repository/dao"
	"github.com/lcsin/webook/internal/service"
	"github.com/lcsin/webook/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		//ioc.InitDB, ioc.InitRedis, handler.InitHandlers, handler.InitMiddlewares, ioc.InitRouter,
		//dao.NewUserDAO, cache.NewUserCache, repository.NewUserRepository, service.NewUserService, handler.NewUserHandler,
		//dao.NewArticleDAO, repository.NewArticleRepository, service.NewArticleService, handler.NewArticleHandler,

		thirdProvider, userHandlerProvider, articleHandlerProvider,
	)

	return gin.Default()
}

// 第三方依赖注入
var thirdProvider = wire.NewSet(
	ioc.InitDB, ioc.InitRedis, ioc.InitRouter,
	handler.InitHandlers, handler.InitMiddlewares,
)

// userHandler依赖注入
var userHandlerProvider = wire.NewSet(
	dao.NewUserDAO,
	cache.NewUserCache,
	repository.NewUserRepository,
	service.NewUserService,
	handler.NewUserHandler,
)

// articleHandler依赖注入
var articleHandlerProvider = wire.NewSet(
	dao.NewArticleDAO,
	repository.NewArticleRepository,
	service.NewArticleService,
	handler.NewArticleHandler,
)

// 初始化测试用的articleHandler
func InitTestArticleHandler() *handler.ArticleHandler {
	wire.Build(
		ioc.InitTestDB, dao.NewUserDAO, articleHandlerProvider,
	)
	return new(handler.ArticleHandler)
}
