package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	iamv1 "github.com/ming-0x0/yuan/pkg/proto/iam/v1"
	blogv1 "github.com/ming-0x0/yuan/pkg/proto/blog/v1"
	"google.golang.org/grpc"
)

type GRPCGateway struct {
	mux     *runtime.ServeMux
	server  *http.Server
}

func NewGRPCGateway(grpcAddr string, grpcConn *grpc.ClientConn) (*GRPCGateway, error) {
	mux := runtime.NewServeMux()

	// Register IAM services
	if err := iamv1.RegisterAuthServiceHandler(context.Background(), mux, grpcConn); err != nil {
		return nil, fmt.Errorf("failed to register auth service handler: %v", err)
	}
	if err := iamv1.RegisterUserServiceHandler(context.Background(), mux, grpcConn); err != nil {
		return nil, fmt.Errorf("failed to register user service handler: %v", err)
	}

	// Register Blog service
	if err := blogv1.RegisterBlogServiceHandler(context.Background(), mux, grpcConn); err != nil {
		return nil, fmt.Errorf("failed to register blog service handler: %v", err)
	}

	return &GRPCGateway{
		mux: mux,
	}, nil
}

func (g *GRPCGateway) Start(port int) error {
	addr := fmt.Sprintf(":%d", port)
	g.server = &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Trim trailing slash
			r.URL.Path = strings.TrimRight(r.URL.Path, "/")

			// Serve the request
			g.mux.ServeHTTP(w, r)
		}),
	}

	fmt.Printf("gRPC gateway listening on port %d\n", port)
	if err := g.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to serve gateway: %v", err)
	}

	return nil
}

func (g *GRPCGateway) Stop(ctx context.Context) error {
	return g.server.Shutdown(ctx)
}
