package logger

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
)

type combinedLogger struct {
	otelLogger *slog.Logger
}

func NewOTelLogger(serviceName, version string) Logger {
	otelLogger := otelslog.NewLogger(serviceName,
		otelslog.WithVersion(version),
		otelslog.WithSource(true),
	)

	return &combinedLogger{
		otelLogger: otelLogger,
	}
}

func (l *combinedLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.otelLogger.DebugContext(ctx, msg, keysAndValues...)
}

func (l *combinedLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.otelLogger.InfoContext(ctx, msg, keysAndValues...)
}

func (l *combinedLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.otelLogger.WarnContext(ctx, msg, keysAndValues...)
}

func (l *combinedLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.otelLogger.ErrorContext(ctx, msg, keysAndValues...)
}

func (l *combinedLogger) Fatal(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.otelLogger.ErrorContext(ctx, msg, keysAndValues...)
	os.Exit(1)
}
