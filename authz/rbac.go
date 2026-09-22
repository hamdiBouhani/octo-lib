package authz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role")
		if userRole != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}
