package grpc

import (
	"context"

	"github.com/ming-0x0/yuan/internal/iam/application"
)

type IAMHandler struct {
	service application.IAMService
}

func NewIAMHandler(service application.IAMService) *IAMHandler {
	return &IAMHandler{service: service}
}

func (h *IAMHandler) Register(ctx context.Context, req application.RegisterRequest) (string, error) {
	userId, err := h.service.Register(ctx, req)
	if err != nil {
		return "", err
	}
	// Fixed: Use String() method on id.ID
	return userId.String(), nil
}

func (h *IAMHandler) Login(ctx context.Context, username, password string) (string, error) {
	token, err := h.service.Login(ctx, username, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
