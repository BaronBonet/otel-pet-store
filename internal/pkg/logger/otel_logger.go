package logger

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
)

type logger struct {
	otelLogger *slog.Logger
}

func NewOTelLogger(serviceName, version string) Logger {
	otelLogger := otelslog.NewLogger(serviceName,
		otelslog.WithVersion(version),
		otelslog.WithSource(true),
	)

	return &logger{
		otelLogger: otelLogger,
	}
}

// logWithLevel is a helper function that handles the common logging logic
// It creates a record with the correct caller program counter and adds all key-value pairs
func (l *logger) logWithLevel(ctx context.Context, level slog.Level, msg string, keysAndValues ...interface{}) {
	// Get the caller's program counter to correctly identify the source location
	// Skipping 2 levels because:
	// - Level 0 points to logWithLevel
	// - Level 1 points to the log function that called logWithLevel
	// - Level 2 points to the actual application code that called the log function
	pc, _, _, _ := runtime.Caller(2)

	r := slog.NewRecord(time.Now(), level, msg, pc)

	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			key, ok := keysAndValues[i].(string)
			if !ok {
				key = "INVALID_KEY"
			}
			value := keysAndValues[i+1]
			r.AddAttrs(slog.Any(key, value))
		}
	}

	_ = l.otelLogger.Handler().Handle(ctx, r)
}

func (l *logger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logWithLevel(ctx, slog.LevelDebug, msg, keysAndValues...)
}

func (l *logger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logWithLevel(ctx, slog.LevelInfo, msg, keysAndValues...)
}

func (l *logger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logWithLevel(ctx, slog.LevelWarn, msg, keysAndValues...)
}

func (l *logger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logWithLevel(ctx, slog.LevelError, msg, keysAndValues...)
}

func (l *logger) Fatal(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logWithLevel(ctx, slog.LevelError, msg, keysAndValues...)
	os.Exit(1)
}
