package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/config"
	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/handler/middleware"
)

type IHandler interface {
	RegisterRoutes(v1 *gin.Engine)
}

func InitHandlers(userHandler *UserHandler, artHandler *ArticleHandler) []IHandler {
	return []IHandler{
		userHandler, artHandler,
	}
}

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.DefaultCors(),
		middleware.NewLoginValidatorBuilder("uid").
			IgnorePath("/users/v1/register").
			IgnorePath("/users/v1/login").
			JWT([]byte(config.Cfg.Jwt.Key), &domain.UserClaims{}),
	}
}
