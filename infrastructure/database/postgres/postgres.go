package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	dsn             string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime int
	connMaxIdleTime int
}

func WithDSN(host string, port int, user string, password string, dbname string, sslMode string) Option {
	return func(cfg *config) {
		cfg.dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, strconv.Itoa(port), user, password, dbname, sslMode)
	}
}

func WithMaxOpenConns(maxOpenConns int) Option {
	return func(cfg *config) {
		cfg.maxOpenConns = maxOpenConns
	}
}

func WithMaxIdleConns(maxIdleConns int) Option {
	return func(cfg *config) {
		cfg.maxIdleConns = maxIdleConns
	}
}

func WithConnMaxLifetime(connMaxLifetime int) Option {
	return func(cfg *config) {
		cfg.connMaxLifetime = connMaxLifetime
	}
}

func WithConnMaxIdleTime(connMaxIdleTime int) Option {
	return func(cfg *config) {
		cfg.connMaxIdleTime = connMaxIdleTime
	}
}

type Option func(*config)

func New(opts ...Option) (*sql.DB, error) {
	cfg := &config{
		maxOpenConns:    10,
		maxIdleConns:    10,
		connMaxLifetime: 300,
		connMaxIdleTime: 60,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	// open the database connection pool
	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		return nil, err
	}

	// set the connection pool configuration
	db.SetMaxOpenConns(cfg.maxOpenConns)
	db.SetMaxIdleConns(cfg.maxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.connMaxLifetime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(cfg.connMaxIdleTime) * time.Second)

	// ping the database
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func Close(db *sql.DB) error {
	return db.Close()
}
