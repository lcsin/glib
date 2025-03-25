package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(config cors.Config) gin.HandlerFunc {
	return cors.New(config)
}

func DefaultCors() gin.HandlerFunc {
	const auth = "Authorization"
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, auth)
	config.ExposeHeaders = append(config.ExposeHeaders, auth)
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	return cors.New(config)
}
