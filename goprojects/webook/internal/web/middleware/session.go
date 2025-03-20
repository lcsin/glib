package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/pkg"
)

func Session(ssid string, secret []byte) gin.HandlerFunc {
	store := cookie.NewStore(secret)
	return sessions.Sessions(ssid, store)
}

type LoginBuilder struct {
	Paths []string
}

func NewLoginBuilder() *LoginBuilder {
	return &LoginBuilder{}
}

func (l *LoginBuilder) IgnorePath(path string) *LoginBuilder {
	l.Paths = append(l.Paths, path)
	return l
}

func (l *LoginBuilder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接口放行
		for _, v := range l.Paths {
			if c.Request.URL.Path == v {
				return
			}
		}

		// 登录校验
		session := sessions.Default(c)
		uid := session.Get("uid")
		if uid == nil { // 用户未登录
			pkg.ResponseError(c, -1, "未授权")
			c.Abort()
			return
		}
		// gin上下文注入uid
		c.Set("uid", uid)
	}
}
