package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ming-0x0/yuan/configs"
	blogHttp "github.com/ming-0x0/yuan/internal/blog/adapter/handler/http"
	blogRepo "github.com/ming-0x0/yuan/internal/blog/adapter/persistence/postgres"
	blog "github.com/ming-0x0/yuan/internal/blog/application"
	authHttp "github.com/ming-0x0/yuan/internal/iam/adapter/handler/http"
	accountRepo "github.com/ming-0x0/yuan/internal/iam/adapter/repository/account"
	"github.com/ming-0x0/yuan/internal/iam/application/auth"
	"github.com/ming-0x0/yuan/internal/iam/application/user"
	"github.com/ming-0x0/yuan/internal/server"
	"github.com/ming-0x0/yuan/pkg/infrastructure/database/postgres"
	"github.com/ming-0x0/yuan/pkg/logger/slog"
	"github.com/ming-0x0/yuan/pkg/orm/bun/client"
	"github.com/ming-0x0/yuan/pkg/timezone"
)

func main() {
	timezone.SetTimeZoneUTC()

	// Load Config
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	fmt.Printf("Config loaded: Env=%s\n", cfg.Env)

	// Initialize Database
	db, err := postgres.New(
		postgres.WithDSN(
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.Database,
			"disable", // sslmode
		),
		postgres.WithMaxOpenConns(cfg.Postgres.MaxOpenConns),
		postgres.WithMaxIdleConns(cfg.Postgres.MaxIdleConns),
		postgres.WithConnMaxLifetime(cfg.Postgres.ConnMaxLifetime),
		postgres.WithConnMaxIdleTime(cfg.Postgres.ConnMaxIdleTime),
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer postgres.Close(db)

	// Initialize Logger
	logHelper := slog.New(slog.WithLevel(cfg.Logger.Level))

	// Initialize Database Client
	bunClient, err := client.NewPostgresClient(db)
	if err != nil {
		log.Fatalf("failed to create bun client: %v", err)
	}

	// Initialize Repositories
	accRepo := accountRepo.New(bunClient, logHelper)
	blogRepo := blogRepo.New(bunClient, logHelper)

	// Initialize Services
	authSvc := auth.NewAuthService(accRepo)
	userSvc := user.NewUserService(accRepo)
	blogSvc := blog.NewBlogService(blogRepo)

	// Initialize Handlers
	authHandler := authHttp.NewAuthHandler(authSvc)
	userHandler := authHttp.NewUserHandler(userSvc)
	blogHandler := blogHttp.NewBlogHandler(blogSvc)

	// Initialize Server
	srv := server.NewServer()

	// Register Routes
	apiGroup := srv.Echo().Group("/api/v1")
	authHttp.RegisterHandlers(apiGroup, authHandler)
	authHttp.RegisterUserRoutes(apiGroup, userHandler)
	blogHttp.RegisterHandlers(apiGroup, blogHandler)

	// Start Server
	go func() {
		if err := srv.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("shutting down the server: %v", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Echo().Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
