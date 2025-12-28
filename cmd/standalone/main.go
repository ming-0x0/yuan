package main

import (
	"os"

	"github.com/ming-0x0/yuan/config"
	"github.com/ming-0x0/yuan/infrastructure/database/postgres"
	sloglogger "github.com/ming-0x0/yuan/pkg/logger/slog"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger := sloglogger.New(
		sloglogger.WithLevel(config.Logger.Level),
		sloglogger.WithWriter(os.Stdout),
	)

	db, err := postgres.New(
		postgres.WithDSN(config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.Database, "disable"),
		postgres.WithMaxOpenConns(config.Postgres.MaxOpenConns),
		postgres.WithMaxIdleConns(config.Postgres.MaxIdleConns),
		postgres.WithConnMaxLifetime(config.Postgres.ConnMaxLifetime),
		postgres.WithConnMaxIdleTime(config.Postgres.ConnMaxIdleTime),
	)
	if err != nil {
		logger.Error("Error while opening the PostgreSQL DB connection pool", "error", err)
		return
	}
	logger.Info("PostgreSQL DB connection pool opened successfully")

	defer func() {
		if err := postgres.Close(db); err != nil {
			logger.Error("Error while closing the PostgreSQL DB connection pool", "error", err)
		}
	}()
	logger.Info("PostgreSQL DB connection pool closed successfully")
}
