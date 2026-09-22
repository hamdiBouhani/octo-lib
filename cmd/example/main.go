package main

import (
	"net/http"

	"github.com/hamdiBouhani/octo-lib/config"
	"github.com/hamdiBouhani/octo-lib/db"
	"github.com/hamdiBouhani/octo-lib/httpserver"
	"github.com/hamdiBouhani/octo-lib/internal/health"
	"github.com/hamdiBouhani/octo-lib/logger"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel)
	defer log.Sync()

	dbConn, err := db.Connect(cfg.DBURL)
	if err != nil {
		log.Fatal("db connect failed", zap.Error(err))
	}
	if err := dbConn.Migrate("migrations"); err != nil {
		log.Fatal("db migrate failed", zap.Error(err))
	}

	mux := http.NewServeMux()
	mux.Handle("/health", health.Handler(dbConn.DB))

	srv := httpserver.New(cfg.HTTPPort, mux)

	if err := srv.Start(); err != nil {
		log.Error("server shutdown error", zap.Error(err))
	}
}
