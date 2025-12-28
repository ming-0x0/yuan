package logrus

import (
	"context"
	"fmt"
	"io"
	"time"

	"runtime"

	"github.com/ming-0x0/yuan/pkg/logger"
	"github.com/sirupsen/logrus"
)

type config struct {
	level  logLevel
	writer io.Writer
}

type Option func(*config)

func WithLevel(level string) Option {
	return func(c *config) {
		c.level = logLevel(level)
	}
}

func WithWriter(writer io.Writer) Option {
	return func(c *config) {
		c.writer = writer
	}
}

type logLevel string

const (
	Fatal logLevel = "fatal"
	Error logLevel = "error"
	Warn  logLevel = "warn"
	Info  logLevel = "info"
	Debug logLevel = "debug"
)

func (l logLevel) toLogrusLevel() logrus.Level {
	switch l {
	case Fatal:
		return logrus.FatalLevel
	case Error:
		return logrus.ErrorLevel
	case Warn:
		return logrus.WarnLevel
	case Info:
		return logrus.InfoLevel
	case Debug:
		return logrus.DebugLevel
	default:
		return logrus.InfoLevel
	}
}

type Logger struct {
	entry *logrus.Entry
}

func New(opts ...Option) *Logger {
	cfg := &config{
		level: Info,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	logger := logrus.New()

	logger.SetOutput(cfg.writer)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     true,
	})
	logger.SetLevel(cfg.level.toLogrusLevel())

	return &Logger{entry: logrus.NewEntry(logger)}
}

func (l *Logger) log(ctx context.Context, level logrus.Level, msg string, keyVals ...any) {
	fields := toFields(keyVals...)

	if _, file, line, ok := runtime.Caller(2); ok {
		pc, _, _, _ := runtime.Caller(2)
		fn := runtime.FuncForPC(pc)
		fields["file"] = fmt.Sprintf("%s:%d", file, line)
		if fn != nil {
			fields["func"] = fn.Name()
		}
	}

	l.entry.WithContext(ctx).WithFields(fields).Log(level, msg)
}

func (l *Logger) Fatal(msg string, keyVals ...any) {
	l.log(context.Background(), logrus.FatalLevel, msg, keyVals...)
}

func (l *Logger) FatalContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, logrus.FatalLevel, msg, keyVals...)
}

func (l *Logger) Error(msg string, keyVals ...any) {
	l.log(context.Background(), logrus.ErrorLevel, msg, keyVals...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, logrus.ErrorLevel, msg, keyVals...)
}

func (l *Logger) Warn(msg string, keyVals ...any) {
	l.log(context.Background(), logrus.WarnLevel, msg, keyVals...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, logrus.WarnLevel, msg, keyVals...)
}

func (l *Logger) Info(msg string, keyVals ...any) {
	l.log(context.Background(), logrus.InfoLevel, msg, keyVals...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, logrus.InfoLevel, msg, keyVals...)
}

func (l *Logger) Debug(msg string, keyVals ...any) {
	l.log(context.Background(), logrus.DebugLevel, msg, keyVals...)
}

func (l *Logger) DebugContext(ctx context.Context, msg string, keyVals ...any) {
	l.log(ctx, logrus.DebugLevel, msg, keyVals...)
}

func (l *Logger) WithFields(keyVals ...any) logger.Logger {
	return &Logger{
		entry: l.entry.WithFields(toFields(keyVals...)),
	}
}

func toFields(keyVals ...any) logrus.Fields {
	fields := logrus.Fields{}

	for i := 0; i+1 < len(keyVals); i += 2 {
		key, ok := keyVals[i].(string)
		if !ok {
			key = fmt.Sprintf("key_%d", i)
		}
		fields[key] = keyVals[i+1]
	}

	return fields
}
