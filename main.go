package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/ming-0x0/yuan/configs"
	blogHttp "github.com/ming-0x0/yuan/internal/blog/adapter/handler/http"
	blogGrpc "github.com/ming-0x0/yuan/internal/blog/adapter/handler/grpc"
	blogRepo "github.com/ming-0x0/yuan/internal/blog/adapter/persistence/postgres"
	blog "github.com/ming-0x0/yuan/internal/blog/application"
	authHttp "github.com/ming-0x0/yuan/internal/iam/adapter/handler/http"
	authGrpc "github.com/ming-0x0/yuan/internal/iam/adapter/handler/grpc"
	userGrpc "github.com/ming-0x0/yuan/internal/iam/adapter/handler/grpc"
	accountRepo "github.com/ming-0x0/yuan/internal/iam/adapter/repository/account"
	"github.com/ming-0x0/yuan/internal/iam/application/auth"
	"github.com/ming-0x0/yuan/internal/iam/application/user"
	"github.com/ming-0x0/yuan/internal/server"
	"github.com/ming-0x0/yuan/pkg/infrastructure/database/postgres"
	"github.com/ming-0x0/yuan/pkg/logger/slog"
	"github.com/ming-0x0/yuan/pkg/orm/bun/client"
	"github.com/ming-0x0/yuan/pkg/timezone"
	"google.golang.org/grpc"
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

	// Initialize HTTP Handlers
	authHandler := authHttp.NewAuthHandler(authSvc)
	userHandler := authHttp.NewUserHandler(userSvc)
	blogHandler := blogHttp.NewBlogHandler(blogSvc)

	// Initialize HTTP Server
	httpSrv := server.NewServer()

	// Register Routes
	apiGroup := httpSrv.Echo().Group("/api/v1")
	authHttp.RegisterHandlers(apiGroup, authHandler)
	authHttp.RegisterUserRoutes(apiGroup, userHandler)
	blogHttp.RegisterHandlers(apiGroup, blogHandler)

	// Initialize gRPC Handlers
	authGrpcHandler := authGrpc.NewAuthServer(authSvc)
	userGrpcHandler := userGrpc.NewUserServer(userSvc)
	blogGrpcHandler := blogGrpc.NewBlogServer(blogSvc)

	// Initialize gRPC Server
	grpcSrv := server.NewGRPCServer(authGrpcHandler, userGrpcHandler, blogGrpcHandler)

	// Initialize gRPC Gateway
	grpcAddr := "localhost:9090"
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	gateway, err := server.NewGRPCGateway(grpcAddr, conn)
	if err != nil {
		log.Fatalf("failed to create gRPC gateway: %v", err)
	}

	// Start Servers
	go func() {
		if err := httpSrv.Start(":8080"); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	go func() {
		if err := grpcSrv.Start(9090); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	go func() {
		if err := gateway.Start(8081); err != nil {
			log.Fatalf("gRPC gateway error: %v", err)
		}
	}()

	fmt.Println("Servers started successfully:")
	fmt.Printf("- HTTP server: http://localhost:8080\n")
	fmt.Printf("- gRPC server: grpc://localhost:9090\n")
	fmt.Printf("- gRPC gateway: http://localhost:8081\n")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("\nShutting down servers...")

	// Shutdown HTTP server
	if err := httpSrv.Echo().Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server shutdown error: %v", err)
	}

	// Shutdown gRPC gateway
	if err := gateway.Stop(ctx); err != nil {
		log.Fatalf("gRPC gateway shutdown error: %v", err)
	}

	// Shutdown gRPC server
	if err := grpcSrv.Stop(ctx); err != nil {
		log.Fatalf("gRPC server shutdown error: %v", err)
	}

	fmt.Println("All servers shutdown successfully")
}
