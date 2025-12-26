package bunlogger

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ming-0x0/yuan/pkg/logger"
	"github.com/uptrace/bun"
)

type Logger struct {
	driver            string
	logger            logger.Logger
	slowThreshold     time.Duration
	ignoreNoRowsError bool
	level             string
}

type Option func(*Logger)

// WithDriver sets the driver for the logger (slog or logrus)
func WithDriver(driver string) Option {
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
		l.level = level
	}
}

// New creates a new bun logger with default options (driver: slog, slowThreshold: 0, ignoreNoRowsError: false)
func New(opts ...Option) *Logger {
	logger := &Logger{
		driver:            "slog",
		ignoreNoRowsError: false,
	}
	for _, opt := range opts {
		opt(logger)
	}

	return logger
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

	keyVals = append(keyVals, "orm", "bun")

	switch {
	case event.Err != nil && (!errors.Is(event.Err, sql.ErrNoRows) || !l.ignoreNoRowsError):
		l.logger.ErrorContext(ctx, "SQL Query failed", keyVals...)
	case l.slowThreshold != 0 && elapsed > l.slowThreshold:
		l.logger.WarnContext(ctx, "Performed SLOW SQL Query", keyVals...)
	default:
		l.logger.InfoContext(ctx, "Performed SQL Query", keyVals...)
	}
}
