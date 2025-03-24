package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lcsin/webook/pkg"
)

type LoginValidatorBuilder struct {
	Key   string
	Paths []string
}

func NewLoginValidatorBuilder(key string) *LoginValidatorBuilder {
	return &LoginValidatorBuilder{
		Key: key,
	}
}

func (l *LoginValidatorBuilder) IgnorePath(path string) *LoginValidatorBuilder {
	l.Paths = append(l.Paths, path)
	return l
}

// 校验接口放行
func (l *LoginValidatorBuilder) checkIgnorePaths(c *gin.Context) bool {
	for _, v := range l.Paths {
		if c.Request.URL.Path == v {
			return true
		}
	}
	return false
}

// Session 通过Session校验用户是否登录
func (l *LoginValidatorBuilder) Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接口放行
		if l.checkIgnorePaths(c) {
			return
		}

		// 登录校验
		session := sessions.Default(c)
		user := session.Get(l.Key)
		if user == nil {
			pkg.ResponseError(c, -1, "未授权")
			c.Abort()
			return
		}

		// gin上下文注入uid
		c.Set(l.Key, user)
	}
}

// JWT 通过JWT校验用户是否登录
func (l *LoginValidatorBuilder) JWT(secret []byte, claims jwt.Claims) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接口放行
		if l.checkIgnorePaths(c) {
			return
		}

		// 从请求头中获取jwt
		hToken := c.GetHeader("Authorization")
		if hToken == "" {
			pkg.ResponseError(c, -1, "未授权")
			c.Abort()
			return
		}

		// 校验jwt
		token, err := jwt.ParseWithClaims(hToken, claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil {
			pkg.ResponseError(c, -1, "未授权")
			c.Abort()
			return
		}
		if token == nil || !token.Valid || claims == nil {
			pkg.ResponseError(c, -1, "未授权")
			c.Abort()
			return
		}

		// gin上下文注入uid
		c.Set(l.Key, claims)
	}
}
