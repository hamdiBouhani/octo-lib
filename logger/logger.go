// logger/logger.go
package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func New(level string) *zap.Logger {
    cfg := zap.NewProductionConfig()
    if err := cfg.Level.UnmarshalText([]byte(level)); err != nil {
        cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    }
    log, _ := cfg.Build()
    return log
}
