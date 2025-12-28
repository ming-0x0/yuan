package slog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/ming-0x0/yuan/pkg/logger"
	"github.com/ming-0x0/yuan/pkg/logger/slog/handler"
)

type config struct {
	level  string
	writer io.Writer
}

type Option func(*config)

func WithLevel(level string) Option {
	return func(c *config) {
		c.level = level
	}
}

func WithWriter(writer io.Writer) Option {
	return func(c *config) {
		c.writer = writer
	}
}

const (
	Fatal = slog.Level(12)
	Error = slog.Level(8)
	Warn  = slog.Level(4)
	Info  = slog.Level(0)
	Debug = slog.Level(-4)
)

var levelNames = map[slog.Leveler]string{
	Fatal: "fatal",
	Error: "error",
	Warn:  "warn",
	Info:  "info",
	Debug: "debug",
}

func getLogLevel(logLevel string) slog.Level {
	switch logLevel {
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
	default:
		return Info
	}
}

type Logger struct {
	*slog.Logger
}

func New(opts ...Option) *Logger {
	cfg := &config{
		level: "info",
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
	fields := toFields(keyVals...)

	if _, file, line, ok := runtime.Caller(2); ok {
		pc, _, _, _ := runtime.Caller(2)
		fn := runtime.FuncForPC(pc)
		fields = append(fields, slog.String("file", fmt.Sprintf("%s:%d", file, line)))
		if fn != nil {
			fields = append(fields, slog.String("func", fn.Name()))
		}
	}

	l.Logger.Log(ctx, level, msg, fields...)
}

// Fatal logs a message with level Fatal on the logger then calls os.Exit(1)
func (l *Logger) Fatal(msg string, attrs ...any) {
	l.log(context.Background(), Fatal, msg, attrs...)
	os.Exit(1)
}

// FatalContext logs a message with level Fatal on the logger then calls os.Exit(1)
func (l *Logger) FatalContext(ctx context.Context, msg string, attrs ...any) {
	l.log(ctx, Fatal, msg, attrs...)
	os.Exit(1)
}

func (l *Logger) Error(msg string, attrs ...any) {
	l.log(context.Background(), Error, msg, attrs...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, attrs ...any) {
	l.log(ctx, Error, msg, attrs...)
}

func (l *Logger) Warn(msg string, attrs ...any) {
	l.log(context.Background(), Warn, msg, attrs...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, attrs ...any) {
	l.log(ctx, Warn, msg, attrs...)
}

func (l *Logger) Info(msg string, attrs ...any) {
	l.log(context.Background(), Info, msg, attrs...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, attrs ...any) {
	l.log(ctx, Info, msg, attrs...)
}

func (l *Logger) Debug(msg string, attrs ...any) {
	l.log(context.Background(), Debug, msg, attrs...)
}

func (l *Logger) DebugContext(ctx context.Context, msg string, attrs ...any) {
	l.log(ctx, Debug, msg, attrs...)
}

// With returns a Logger that includes the given attributes in each output operation.
func (l *Logger) WithFields(args ...any) logger.Logger {
	return &Logger{Logger: l.Logger.With(args...)}
}

func toFields(keyVals ...any) []any {
	fields := make([]any, 0, len(keyVals)/2)

	for i := 0; i+1 < len(keyVals); i += 2 {
		key, ok := keyVals[i].(string)
		if !ok {
			key = fmt.Sprintf("key_%d", i)
		}
		fields = append(fields, slog.Attr{
			Key:   key,
			Value: slog.AnyValue(keyVals[i+1]),
		})
	}

	return fields
}
