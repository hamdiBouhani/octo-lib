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
