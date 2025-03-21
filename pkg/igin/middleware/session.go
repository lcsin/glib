package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Session(ssid string, store sessions.Store) gin.HandlerFunc {
	return sessions.Sessions(ssid, store)
}
