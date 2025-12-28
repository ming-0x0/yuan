package bunlogger

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ming-0x0/yuan/pkg/logger"
	sloglogger "github.com/ming-0x0/yuan/pkg/logger/slog"
	"github.com/uptrace/bun"
)

type Driver string

const (
	Slog   Driver = "slog"
	Logrus Driver = "logrus"
)

type Logger struct {
	driver            Driver
	logger            logger.Logger
	slowThreshold     time.Duration
	ignoreNoRowsError bool
	level             logger.Level
}

type Option func(*Logger)

// WithDriver sets the driver for the logger (slog or logrus)
func WithDriver(driver Driver) Option {
	return func(l *Logger) {
		l.driver = driver
	}
}

// WithLogger sets the logger for the bun logger
func WithLogger(logger logger.Logger) Option {
	return func(l *Logger) {
		l.logger = logger
	}
}

// WithSlowThreshold sets the slow threshold for the bun logger
func WithSlowThreshold(slowThreshold time.Duration) Option {
	return func(l *Logger) {
		l.slowThreshold = slowThreshold
	}
}

// WithIgnoreNoRowsError sets the ignore no rows error for the bun logger
func WithIgnoreNoRowsError() Option {
	return func(l *Logger) {
		l.ignoreNoRowsError = true
	}
}

// WithLevel sets the level for the bun logger
func WithLevel(level string) Option {
	return func(l *Logger) {
		l.level = logger.ParseLevel(level)
	}
}

// New creates a new bun logger with default options (driver: slog, slowThreshold: 0, ignoreNoRowsError: false)
func New(opts ...Option) *Logger {
	l := &Logger{
		driver:            Slog,
		ignoreNoRowsError: false,
		level:             logger.Info,
	}
	for _, opt := range opts {
		opt(l)
	}

	if l.logger == nil {
		if l.driver == Slog {
			l.logger = sloglogger.New()
		}
	}

	return l
}

var _ bun.QueryHook = (*Logger)(nil)

func (l *Logger) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (l *Logger) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	elapsed := time.Since(event.StartTime)

	// Build attributes
	keyVals := []any{
		"durations", elapsed.String(),
		"sql", event.Query,
	}

	if event.Result != nil {
		rows, err := event.Result.RowsAffected()
		if err != nil {
			keyVals = append(keyVals, "rows", "-")
		} else {
			keyVals = append(keyVals, "rows", rows)
		}
	}

	if event.Err != nil {
		keyVals = append(keyVals, "error", event.Err.Error())
	}

	keyVals = append(keyVals, "service", "database")

	switch {
	case event.Err != nil && (!errors.Is(event.Err, sql.ErrNoRows) || !l.ignoreNoRowsError):
		if l.level.Priority() >= logger.Error.Priority() {
			l.logger.ErrorContext(ctx, "SQL Query failed", keyVals...)
		}
	case l.slowThreshold != 0 && elapsed > l.slowThreshold:
		if l.level.Priority() >= logger.Warn.Priority() {
			l.logger.WarnContext(ctx, "Performed SLOW SQL Query", keyVals...)
		}
	default:
		if l.level.Priority() >= logger.Info.Priority() {
			l.logger.InfoContext(ctx, "Performed SQL Query", keyVals...)
		}
	}
}
