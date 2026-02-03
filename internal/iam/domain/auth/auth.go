package auth

import (
	"context"
)

type AuthService interface {
	Register(ctx context.Context, email string, password string) error
}
