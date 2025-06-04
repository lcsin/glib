package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/glib/pkg/igin/middleware"
)

type IHandler interface {
	RegisterRoutes(r *gin.Engine)
}

func InitHandlers(bookHandler *BookHandler) []IHandler {
	return []IHandler{
		bookHandler,
	}
}

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.DefaultCors(),
	}
}
