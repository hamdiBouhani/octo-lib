package db

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm/logger"
)

type CtxLogger struct {
	logLevel logger.LogLevel
}

func NewCtxLogger() *CtxLogger {
	return &CtxLogger{
		logLevel: logger.Info,
	}
}

func (l *CtxLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.logLevel = level
	return l
}

func (l *CtxLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.print(ctx, "INFO", msg, data...)
}

func (l *CtxLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.print(ctx, "WARN", msg, data...)
}

func (l *CtxLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.print(ctx, "ERROR", msg, data...)
}

func (l *CtxLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)

	if err != nil {
		l.print(ctx, "ERROR", err.Error(), sql, rows, elapsed)
		return
	}

	l.print(ctx, "SQL", sql, rows, elapsed)
}

func (l *CtxLogger) print(ctx context.Context, level string, msg string, data ...interface{}) {
	reqID := ctx.Value("request_id")
	traceID := ctx.Value("trace_id")

	prefix := ""

	if reqID != nil {
		prefix += "[request_id=" + reqID.(string) + "] "
	}

	if traceID != nil {
		prefix += "[trace_id=" + traceID.(string) + "] "
	}

	log.Printf("%s%s %v", prefix, level, append([]interface{}{msg}, data...))
}
