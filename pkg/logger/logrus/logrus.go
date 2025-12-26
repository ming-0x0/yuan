package logrus

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ming-0x0/yuan/pkg/logger"
	"github.com/sirupsen/logrus"
)

type config struct {
	level logLevel
}

type Option func(*config)

func WithLevel(level string) Option {
	return func(c *config) {
		c.level = logLevel(level)
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

	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     true,
	})
	logger.SetLevel(cfg.level.toLogrusLevel())

	return &Logger{entry: logrus.NewEntry(logger)}
}

func (l *Logger) Fatal(msg string, keyVals ...any) {
	l.entry.WithFields(toFields(keyVals...)).Fatal(msg)
}

func (l *Logger) FatalContext(ctx context.Context, msg string, keyVals ...any) {
	l.entry.WithContext(ctx).WithFields(toFields(keyVals...)).Fatal(msg)
}

func (l *Logger) Error(msg string, keyVals ...any) {
	l.entry.WithFields(toFields(keyVals...)).Error(msg)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, keyVals ...any) {
	l.entry.WithContext(ctx).WithFields(toFields(keyVals...)).Error(msg)
}

func (l *Logger) Warn(msg string, keyVals ...any) {
	l.entry.WithFields(toFields(keyVals...)).Warn(msg)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, keyVals ...any) {
	l.entry.WithContext(ctx).WithFields(toFields(keyVals...)).Warn(msg)
}

func (l *Logger) Info(msg string, keyVals ...any) {
	l.entry.WithFields(toFields(keyVals...)).Info(msg)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, keyVals ...any) {
	l.entry.WithContext(ctx).WithFields(toFields(keyVals...)).Info(msg)
}

func (l *Logger) Debug(msg string, keyVals ...any) {
	l.entry.WithFields(toFields(keyVals...)).Debug(msg)
}

func (l *Logger) DebugContext(ctx context.Context, msg string, keyVals ...any) {
	l.entry.WithContext(ctx).WithFields(toFields(keyVals...)).Debug(msg)
}

func (l *Logger) With(keyVals ...any) logger.Logger {
	return &Logger{
		entry: l.entry.WithFields(toFields(keyVals...)),
	}
}

func (l *Logger) WithGroup(name string) logger.Logger {
	return &Logger{
		entry: l.entry.WithFields(logrus.Fields{name: name}),
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
