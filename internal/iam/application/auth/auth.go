package auth

import (
	"context"

	"github.com/ming-0x0/yuan/internal/iam/domain/account"
)

type AuthService interface {
	Register(ctx context.Context, email string, password string, fullName string) error
	Login(ctx context.Context, email string, password string) (string, error)
}

type authService struct {
	accountRepo account.AccountRepository
}

func NewAuthService(accountRepo account.AccountRepository) AuthService {
	return &authService{
		accountRepo: accountRepo,
	}
}
