package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Session() gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	return sessions.Sessions("session", store)
}
