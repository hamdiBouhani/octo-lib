# octo-lib

A shared enterprise-grade Go library providing reusable building blocks for all backend services in our ecosystem.  
It standardizes configuration, logging, HTTP servers, middleware, tracing, metrics, database access, Kafka integration, OAuth2 authentication, RBAC authorization, health checks, utilities, and test helpers.

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
