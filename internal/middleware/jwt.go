package middleware

import (
	"strings"

	"github.com/Chihaya-Anon123/task_manager/internal/code"
	"github.com/Chihaya-Anon123/task_manager/internal/config"
	"github.com/Chihaya-Anon123/task_manager/internal/response"
	"github.com/Chihaya-Anon123/task_manager/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	CtxUserIDKey   = "user_id"
	CtxUsernameKey = "username"
)

func JWTAuth(cfg config.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Fail(c, code.CodeUnauthorized, "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Fail(c, code.CodeUnauthorized, "invalid authorization header")
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(parts[1])
		if tokenString == "" {
			response.Fail(c, code.CodeUnauthorized, "invalid token")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(tokenString, cfg.Secret)
		if err != nil {
			response.Fail(c, code.CodeUnauthorized, "invalid or expired token")
			c.Abort()
			return
		}

		c.Set(CtxUserIDKey, claims.UserID)
		c.Set(CtxUsernameKey, claims.Username)

		c.Next()
	}
}

func GetCurrentUserID(c *gin.Context) (uint, bool) {
	v, ok := c.Get(CtxUserIDKey)
	if !ok {
		return 0, false
	}

	userID, ok := v.(uint)
	return userID, ok
}

func GetCurrentUsername(c *gin.Context) (string, bool) {
	v, ok := c.Get(CtxUsernameKey)
	if !ok {
		return "", false
	}

	username, ok := v.(string)
	return username, ok
}
