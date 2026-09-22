# octo-lib

A shared enterprise-grade Go library providing reusable building blocks for all backend services in our ecosystem.  
It standardizes configuration, logging, HTTP servers, middleware, tracing, metrics, database access, OAuth2 authentication, RBAC authorization, health checks, utilities, and test helpers.

This library ensures consistency, reliability, and observability across all microservices.

---

## 📦 Modules

### **config/**
Environment-based configuration loader using `.env` files.

- Loads `.env` automatically
- Provides required and default values
- Panics on missing required env vars

### **logger/**
Zap-based structured logger.

- Production-ready logging
- JSON output
- Integrates with request_id and trace_id

### **httpserver/**
Gin-based HTTP server wrapper.

Includes:
- Logging middleware
- Recovery middleware
- Request ID middleware

### **middleware/**
Shared middleware for all services.

- Tracing (OpenTelemetry)
- Metrics (Prometheus)
- RBAC authorization

### **tracing/**
OpenTelemetry initialization and helpers.

### **metrics/**
Prometheus registry and HTTP metrics middleware.

### **db/**
GORM database wrapper with contextual logging.

- SQL logs include `request_id` and `trace_id`
- Supports `db.WithContext(ctx)`


### **oauth2/**
JWT/OAuth2 validator and Gin middleware.

- Validates tokens
- Injects user identity into request context

### **authz/**
Simple RBAC authorization middleware.

### **errors/**
Common error definitions.

### **health/**
Standard health endpoint handler.

### **utils/**
Small reusable helpers (UUID, time, env).

### **testutils/**
Godog test helpers for E2E testing.

---

## 🚀 Getting Started

### Install

```bash
go get github.com/hamdiBouhani/octo-lib
```
Example:

```go
package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hamdiBouhani/octo-lib/config"
	"github.com/hamdiBouhani/octo-lib/db"
	"github.com/hamdiBouhani/octo-lib/httpserver"
	"github.com/hamdiBouhani/octo-lib/logger"
	"github.com/hamdiBouhani/octo-lib/tracing"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)
	defer log.Sync()

	// Init tracing
	tp, err := tracing.Init("octo-lib", "localhost:4318")
	if err != nil {
		log.Fatal("failed to init tracing", zap.Error(err))
	}
	defer tp.Shutdown(context.Background())

	dbConn, err := db.Connect(cfg.DBURL)
	if err != nil {
		log.Fatal("db connect failed", zap.Error(err))
	}

	srv := httpserver.New(cfg.HTTPPort, log, "octo-lib")
	r := srv.Engine()

	r.GET("/health", func(c *gin.Context) {
		if err := dbConn.Ping(); err != nil {
			c.JSON(503, gin.H{"status": "unhealthy"})
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	})

	srv.Start()
}
```