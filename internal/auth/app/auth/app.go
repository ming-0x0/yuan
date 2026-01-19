package auth

import (
	"github.com/ming-0x0/yuan/internal/auth/domain/account"
	"github.com/ming-0x0/yuan/pkg/logger"
)

type AuthApp struct {
	accountRepo account.AccountRepository
	logger      *logger.Logger
}

func New(
	accountRepo account.AccountRepository,
	logger *logger.Logger,
) *AuthApp {
	return &AuthApp{
		accountRepo: accountRepo,
		logger:      logger,
	}
}
