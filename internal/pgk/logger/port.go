package logger

import "context"

type Logger interface {
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	// Fatal Has a side effect of os.Exit(1)
	Fatal(ctx context.Context, msg string, keysAndValues ...interface{})
}
