package httpserver

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func TracingMiddleware(serviceName string) gin.HandlerFunc {
	tracer := otel.Tracer(serviceName)

	return func(c *gin.Context) {
		ctx, span := tracer.Start(c.Request.Context(), c.FullPath())

		traceID := span.SpanContext().TraceID().String()
		c.Set("trace_id", traceID)

		// If request_id not set, use trace_id
		if c.GetString("request_id") == "" {
			c.Set("request_id", traceID)
		}

		span.SetAttributes(
			attribute.String("http.method", c.Request.Method),
			attribute.String("http.path", c.Request.URL.Path),
			attribute.String("client.ip", c.ClientIP()),
		)

		// Inject context into Gin
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		span.SetAttributes(
			attribute.Int("http.status", c.Writer.Status()),
		)

		span.End()
	}
}
