package main

import (
	"github.com/ming-0x0/yuan/adapter/persistence/postgres/account"
	"github.com/ming-0x0/yuan/adapter/persistence/postgres/client"
	restadapter "github.com/ming-0x0/yuan/adapter/rest"
	accounthandler "github.com/ming-0x0/yuan/adapter/rest/account"
	"github.com/ming-0x0/yuan/config"
	"github.com/ming-0x0/yuan/infrastructure/database/postgres"
	accountapp "github.com/ming-0x0/yuan/internal/account/app/account"
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

	// Adapters
	pgClient := client.New(pgDB)
	accountRepo := account.New(pgClient, logger)

	// App Services
	accountApp := accountapp.New(accountRepo)

	// Handlers
	accountHandler := accounthandler.NewHandler(accountApp)

	// Router
	router := restadapter.NewRouter(accountHandler)

	logger.Info("Starting HTTP server...", "port", cfg.HTTPServer.Port)
	if err := router.Run(cfg.HTTPServer.Port); err != nil {
		logger.Fatal("Failed to start HTTP server", "error", err)
	}
}
