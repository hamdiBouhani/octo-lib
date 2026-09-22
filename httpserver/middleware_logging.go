package httpserver

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Extract trace/span IDs from context
		span := trace.SpanFromContext(c.Request.Context())
		spanCtx := span.SpanContext()

		traceID := spanCtx.TraceID().String()
		spanID := spanCtx.SpanID().String()

		c.Next()

		logger.Info("request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
			zap.String("trace_id", traceID),
			zap.String("span_id", spanID),
		)
	}
}
