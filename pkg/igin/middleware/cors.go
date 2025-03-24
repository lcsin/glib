package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const authorization = "Authorization"

// CORS 自定义跨域中间件
func CORS(config cors.Config) gin.HandlerFunc {
	return cors.New(config)
}

// DefaultCors 默认的跨域中间件
func DefaultCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, authorization)
	config.ExposeHeaders = append(config.ExposeHeaders, authorization)
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	return cors.New(config)
}
