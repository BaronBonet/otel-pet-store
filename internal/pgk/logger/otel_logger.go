package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/attribute"
	otellog "go.opentelemetry.io/otel/log"
)

type combinedLogger struct {
	otelLogger *slog.Logger
}

func NewOTelLogger(serviceName, version string) Logger {
	otelLogger := otelslog.NewLogger(serviceName,
		otelslog.WithVersion(version),
		// TODO: does this work or just show the location of this file?
		otelslog.WithSource(true),
	)

	return &combinedLogger{
		otelLogger: otelLogger,
	}
}

func (l *combinedLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	fmt.Printf("combinedLogger.Debug: %s", msg)
	l.otelLogger.DebugContext(ctx, msg, keysAndValues...)
}

func (l *combinedLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	fmt.Printf("combinedLogger.info: %s", msg)
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

func TestLogExport(ctx context.Context) {
	// Get the global logger provider
	loggerProvider := otellog.GetLoggerProvider()

	// Create a logger from the provider
	logger := loggerProvider.Logger(
		"test.logger",
		otellog.WithInstrumentationVersion("v0.1.0"),
	)

	// Log a message with attributes
	logger.Info(
		"This is a test log message to verify OTLP export",
		otellog.WithAttributes(
			attribute.String("test.attribute", "test-value"),
			attribute.Int("test.number", 42),
		),
	)

	fmt.Println("Test log message sent to exporter")
}
