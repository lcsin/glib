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
		ioc.InitDB, ioc.InitRedis, handler.InitHandlers, handler.InitMiddlewares, ioc.InitRouter,
		dao.NewUserDAO, cache.NewUserCache, repository.NewUserRepository, service.NewUserService, handler.NewUserHandler,
	)

	return gin.Default()
}
