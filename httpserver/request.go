package httpserver

import "github.com/gin-gonic/gin"

func RequestID(c *gin.Context) string {
	return c.GetString("request_id")
}

func TraceID(c *gin.Context) string {
	return c.GetString("trace_id")
}

func UserID(c *gin.Context) string {
	return c.GetString("user_id")
}

func Roles(c *gin.Context) []string {
	if v, ok := c.Get("roles"); ok {
		if roles, ok := v.([]string); ok {
			return roles
		}
	}
	return nil
}
