package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/web"
	"github.com/lcsin/webook/pkg"
)

func Session(name string, store sessions.Store) gin.HandlerFunc {
	return sessions.Sessions(name, store)
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
		user, ok := session.Get(web.SessionUser).(*domain.User)
		if !ok || user == nil { // 用户未登录
			pkg.ResponseError(c, -1, "未授权")
			c.Abort()
			return
		}

		// gin上下文注入uid
		c.Set("uid", user.ID)
	}
}
