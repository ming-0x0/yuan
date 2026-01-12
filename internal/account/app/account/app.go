package app

import (
	"github.com/ming-0x0/yuan/internal/account/domain/account"
)

type AccountApp struct {
	accountRepo account.AccountRepository
}

func New(
	accountRepo account.AccountRepository,
) *AccountApp {
	return &AccountApp{
		accountRepo: accountRepo,
	}
}
