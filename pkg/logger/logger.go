package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/ming-0x0/yuan/pkg/logger/handler"
)

type config struct {
	level  string
	writer io.Writer
}

type option func(*config)

func WithLevel(level string) option {
	return func(c *config) {
		c.level = level
	}
}

func WithWriter(writer io.Writer) option {
	return func(c *config) {
		c.writer = writer
	}
}

const (
	Panic = slog.Level(16)
	Fatal = slog.Level(12)
	Error = slog.Level(8)
	Warn  = slog.Level(4)
	Info  = slog.Level(0)
	Debug = slog.Level(-4)
	Trace = slog.Level(-8)
)

var levelNames = map[slog.Leveler]string{
	Panic: "panic",
	Fatal: "fatal",
	Error: "error",
	Warn:  "warn",
	Info:  "info",
	Debug: "debug",
	Trace: "trace",
}

func getLogLevel(logLevel string) slog.Level {
	switch logLevel {
	case "panic":
		return Panic
	case "fatal":
		return Fatal
	case "error":
		return Error
	case "warn":
		return Warn
	case "info":
		return Info
	case "debug":
		return Debug
	case "trace":
		return Trace
	default:
		return Info
	}
}

type Logger struct {
	*slog.Logger
}

func New(opts ...option) *Logger {
	cfg := &config{
		level:  "info",
		writer: os.Stdout,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	handlerOptions := &slog.HandlerOptions{
		Level: getLogLevel(cfg.level),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				return slog.String(slog.TimeKey, a.Value.Time().Format(time.RFC3339))
			case slog.LevelKey:
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := levelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				return slog.String(slog.LevelKey, levelLabel)
			default:
				return a
			}
		},
	}

	slogHandler := handler.NewJSONHandler(cfg.writer, handlerOptions)

	return &Logger{
		Logger: slog.New(slogHandler),
	}
}

func (l *Logger) log(ctx context.Context, level slog.Level, msg string, keyVals ...any) {
	attrs := toAttrs(keyVals...)

	if _, file, line, ok := runtime.Caller(2); ok {
		pc, _, _, _ := runtime.Caller(2)
		fn := runtime.FuncForPC(pc)
		attrs = append(attrs, slog.String("file", fmt.Sprintf("%s:%d", file, line)))
		if fn != nil {
			attrs = append(attrs, slog.String("func", fn.Name()))
		}
	}

	l.Logger.Log(ctx, level, msg, attrs...)
}

// Panic logs a message with level Panic on the logger then calls panic()
func (l *Logger) Panic(msg string, keyVals ...any) {
	l.log(context.Background(), Panic, msg, keyVals...)
	panic(msg)
}

// PanicContext logs a message with level Panic on the logger then calls panic()
func (l *Logger) PanicContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, Panic, msg, keyVals...)
	panic(msg)
}

// Fatal logs a message with level Fatal on the logger then calls os.Exit(1)
func (l *Logger) Fatal(msg string, keyVals ...any) {
	l.log(context.Background(), Fatal, msg, keyVals...)
	os.Exit(1)
}

// FatalContext logs a message with level Fatal on the logger then calls os.Exit(1)
func (l *Logger) FatalContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, Fatal, msg, keyVals...)
	os.Exit(1)
}

func (l *Logger) Error(msg string, keyVals ...any) {
	l.log(context.Background(), Error, msg, keyVals...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, Error, msg, keyVals...)
}

func (l *Logger) Warn(msg string, keyVals ...any) {
	l.log(context.Background(), Warn, msg, keyVals...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, Warn, msg, keyVals...)
}

func (l *Logger) Info(msg string, keyVals ...any) {
	l.log(context.Background(), Info, msg, keyVals...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, Info, msg, keyVals...)
}

func (l *Logger) Debug(msg string, keyVals ...any) {
	l.log(context.Background(), Debug, msg, keyVals...)
}

func (l *Logger) DebugContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, Debug, msg, keyVals...)
}

func (l *Logger) Trace(msg string, keyVals ...any) {
	l.log(context.Background(), Trace, msg, keyVals...)
}

func (l *Logger) TraceContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, Trace, msg, keyVals...)
}

func toAttrs(keyVals ...any) []any {
	attrs := make([]any, 0, len(keyVals)/2)

	for i := 0; i+1 < len(keyVals); i += 2 {
		key, ok := keyVals[i].(string)
		if !ok {
			key = fmt.Sprintf("key_%d", i)
		}
		attrs = append(attrs, slog.Attr{
			Key:   key,
			Value: slog.AnyValue(keyVals[i+1]),
		})
	}

	return attrs
}
