package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Session(ssid string, secret []byte) gin.HandlerFunc {
	store := cookie.NewStore(secret)
	return sessions.Sessions(ssid, store)
}
