package logger

import "context"

type Level string

const (
	Fatal Level = "fatal"
	Error Level = "error"
	Warn  Level = "warn"
	Info  Level = "info"
	Debug Level = "debug"
)

// Priority returns the numeric priority of the log level.
// Higher values mean more severe levels.
func (l Level) Priority() int {
	switch l {
	case Fatal:
		return 100
	case Error:
		return 80
	case Warn:
		return 60
	case Info:
		return 40
	case Debug:
		return 20
	default:
		return 0
	}
}

// ParseLevel parses a string into a Level.
func ParseLevel(s string) Level {
	switch Level(s) {
	case Fatal, Error, Warn, Info, Debug:
		return Level(s)
	default:
		return Info
	}
}

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
	WithFields(keyVals ...any) Logger
}
