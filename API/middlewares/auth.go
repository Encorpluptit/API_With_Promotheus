package middlewares

import (
	"BackendGo/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := auth.TokenValid(c.Request); err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Next()
	}
}
