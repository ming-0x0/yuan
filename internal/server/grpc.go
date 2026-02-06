package server

import (
	"context"
	"fmt"
	"net"

	iamv1 "github.com/ming-0x0/yuan/pkg/proto/iam/v1"
	blogv1 "github.com/ming-0x0/yuan/pkg/proto/blog/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
}

func NewGRPCServer(
	authHandler iamv1.AuthServiceServer,
	userHandler iamv1.UserServiceServer,
	blogHandler blogv1.BlogServiceServer,
) *GRPCServer {
	s := grpc.NewServer()

	iamv1.RegisterAuthServiceServer(s, authHandler)
	iamv1.RegisterUserServiceServer(s, userHandler)
	blogv1.RegisterBlogServiceServer(s, blogHandler)

	// Enable reflection for debugging
	reflection.Register(s)

	return &GRPCServer{
		server: s,
	}
}

func (s *GRPCServer) Start(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	fmt.Printf("gRPC server listening on port %d\n", port)
	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(done)
	}()

	select {
	case <-ctx.Done():
		s.server.Stop()
		return ctx.Err()
	case <-done:
		return nil
	}
}
