package logger

import "context"

type Logger interface {
	Panic(msg string, keyVals ...any)
	PanicContext(ctx context.Context, msg string, keyVals ...any)
	Fatal(msg string, keyVals ...any)
	FatalContext(ctx context.Context, msg string, keyVals ...any)
	Error(msg string, keyVals ...any)
	ErrorContext(ctx context.Context, msg string, keyVals ...any)
	Warn(msg string, keyVals ...any)
	WarnContext(ctx context.Context, msg string, keyVals ...any)
	Info(msg string, keyVals ...any)
	InfoContext(ctx context.Context, msg string, keyVals ...any)
	Debug(msg string, keyVals ...any)
	DebugContext(ctx context.Context, msg string, keyVals ...any)
	Trace(msg string, keyVals ...any)
	TraceContext(ctx context.Context, msg string, keyVals ...any)
}
