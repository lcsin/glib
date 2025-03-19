package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(config cors.Config) gin.HandlerFunc {
	return cors.New(config)
}

func DefaultCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	return cors.New(config)
}
