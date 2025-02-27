package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS 跨域中间件
func CORS(config cors.Config) gin.HandlerFunc {
	return cors.New(config)
}
