package slog

import (
	"context"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"

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
		Level:     getLogLevel(cfg.level),
		AddSource: true,
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

// Fatal logs a message with level Fatal on the logger then calls os.Exit(1)
func (l *Logger) Fatal(msg string, attrs ...any) {
	r := slog.NewRecord(time.Now(), Fatal, msg, 0)
	// Get the caller's PC (program counter) and file/line info
	var pcs [1]uintptr
	// Skip 2 frames: runtime.Callers and this function
	runtime.Callers(2, pcs[:])
	r.PC = pcs[0]
	_ = l.Handler().Handle(context.Background(), r)
	os.Exit(1)
}

// FatalContext logs a message with level Fatal on the logger then calls os.Exit(1)
func (l *Logger) FatalContext(ctx context.Context, msg string, attrs ...any) {
	r := slog.NewRecord(time.Now(), Fatal, msg, 0)
	// Get the caller's PC (program counter) and file/line info
	var pcs [1]uintptr
	// Skip 2 frames: runtime.Callers and this function
	runtime.Callers(2, pcs[:])
	r.PC = pcs[0]
	_ = l.Handler().Handle(ctx, r)
	os.Exit(1)
}

// Attr returns a new Attr with the given key and value
func (l *Logger) Attr(key string, value any) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.AnyValue(value),
	}
}
