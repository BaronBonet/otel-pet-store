package logger

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

type slogLogger struct {
	logger *slog.Logger
}

// NewSlogLogger creates a new slog logger with the specified log level.
// If no log level is provided (slog.LevelDebug), it defaults to debug.
func NewSlogLogger(logLevel ...slog.Level) Logger {
	level := slog.LevelInfo

	if len(logLevel) > 0 {
		level = logLevel[0]
	}

	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      level,
		TimeFormat: time.TimeOnly,
	}))

	slog.SetDefault(logger)

	return &slogLogger{
		logger: logger,
	}
}

func (l *slogLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logger.Debug(msg, keysAndValues...)
}

func (l *slogLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, keysAndValues...)
}

func (l *slogLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logger.Warn(msg, keysAndValues...)
}

func (l *slogLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, keysAndValues...)
}

func (l *slogLogger) Fatal(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, keysAndValues...)
	os.Exit(1)
}
