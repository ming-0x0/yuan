package main

import (
	"github.com/gin-gonic/gin"
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

	// Dependency Injection with Wire
	handler, err := InitializeApp(pgDB)
	if err != nil {
		logger.Fatal("Failed to initialize application", "error", err)
	}

	// Router
	r := gin.Default()
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	logger.Info("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", "error", err)
	}

	defer postgres.Close(pgDB)
}
