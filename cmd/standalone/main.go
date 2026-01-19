package main

import (
	"github.com/ming-0x0/yuan/config"
	"github.com/ming-0x0/yuan/infrastructure/database/postgres"
	"github.com/ming-0x0/yuan/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger := logger.New(
		logger.WithLevel(cfg.Logger.Level),
	)

	logger.Info("Connecting to database...")
	pgDB, err := postgres.New(
		postgres.WithDSN(cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database, "disable"),
		postgres.WithMaxOpenConns(cfg.Postgres.MaxOpenConns),
		postgres.WithMaxIdleConns(cfg.Postgres.MaxIdleConns),
		postgres.WithConnMaxLifetime(cfg.Postgres.ConnMaxLifetime),
		postgres.WithConnMaxIdleTime(cfg.Postgres.ConnMaxIdleTime),
	)
	if err != nil {
		logger.Fatal("Failed to connect to database", "error", err)
	}

	logger.Info("Connected to database successfully.")
	defer postgres.Close(pgDB)
}
