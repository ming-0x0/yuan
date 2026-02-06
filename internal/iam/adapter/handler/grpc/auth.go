package grpc

import (
	"context"

	iamv1 "github.com/ming-0x0/yuan/pkg/proto/iam/v1"
	"github.com/ming-0x0/yuan/internal/iam/application/auth"
)

type AuthServer struct {
	iamv1.UnimplementedAuthServiceServer
	authSvc auth.AuthService
}

func NewAuthServer(authSvc auth.AuthService) iamv1.AuthServiceServer {
	return &AuthServer{
		authSvc: authSvc,
	}
}

func (s *AuthServer) Register(ctx context.Context, req *iamv1.RegisterRequest) (*iamv1.RegisterResponse, error) {
	err := s.authSvc.Register(ctx, req.Email, req.Password, req.FullName)
	if err != nil {
		return nil, err
	}

	return &iamv1.RegisterResponse{
		Message: "User registered successfully",
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *iamv1.LoginRequest) (*iamv1.LoginResponse, error) {
	token, err := s.authSvc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &iamv1.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   3600 * 24, // 24 hours
	}, nil
}
