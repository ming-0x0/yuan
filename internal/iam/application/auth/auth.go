package auth

import (
	"github.com/ming-0x0/yuan/internal/iam/domain/account"
)

type authService struct {
	accountRepo account.AccountRepository
}

func NewAuthService(accountRepo account.AccountRepository) *authService {
	return &authService{
		accountRepo: accountRepo,
	}
}
