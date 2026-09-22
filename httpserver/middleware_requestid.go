package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// All logs generated during that request share the same request_id
// You can filter logs by request_id and see the entire flow
// Debugging becomes 10× easier
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()
		c.Set("request_id", id)
		c.Writer.Header().Set("X-Request-ID", id)
		c.Next()
	}
}
