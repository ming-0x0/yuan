package logger

import "context"

type Level string

const (
	FatalLevel Level = "fatal"
	ErrorLevel Level = "error"
	WarnLevel  Level = "warn"
	InfoLevel  Level = "info"
	DebugLevel Level = "debug"
)

type Logger interface {
	// Fatal logs a message with level Fatal on the logger then calls os.Exit(1)
	Fatal(msg string, keyVals ...any)
	// FatalContext logs a message with level Fatal on the logger then calls os.Exit(1)
	FatalContext(ctx context.Context, msg string, keyVals ...any)
	// Error logs a message with level Error on the logger
	Error(msg string, keyVals ...any)
	// ErrorContext logs a message with level Error on the logger
	ErrorContext(ctx context.Context, msg string, keyVals ...any)
	// Warn logs a message with level Warn on the logger
	Warn(msg string, keyVals ...any)
	// WarnContext logs a message with level Warn on the logger
	WarnContext(ctx context.Context, msg string, keyVals ...any)
	// Info logs a message with level Info on the logger
	Info(msg string, keyVals ...any)
	// InfoContext logs a message with level Info on the logger
	InfoContext(ctx context.Context, msg string, keyVals ...any)
	// Debug logs a message with level Debug on the logger
	Debug(msg string, keyVals ...any)
	// DebugContext logs a message with level Debug on the logger
	DebugContext(ctx context.Context, msg string, keyVals ...any)
	// With returns a Logger that includes the given attributes in each output operation.
	With(keyVals ...any) Logger
}
