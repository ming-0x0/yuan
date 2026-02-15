package auth

import (
	"context"

	"github.com/ming-0x0/yuan/internal/modules/iam/application"
)

type AuthHandler struct {
	service application.IAMService
}

func NewAuthHandler(service application.IAMService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(ctx context.Context, req application.RegisterRequest) (string, error) {
	userId, err := h.service.Register(ctx, req)
	if err != nil {
		return "", err
	}
	return userId.String(), nil
}

func (h *AuthHandler) Login(ctx context.Context, username, password string) (string, error) {
	token, err := h.service.Login(ctx, username, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
