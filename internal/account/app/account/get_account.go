package app

import (
	"context"
	"errors"

	"github.com/ming-0x0/yuan/internal/account/domain/account"
	"github.com/ming-0x0/yuan/internal/common/apperror"
)

func (a *AccountApp) GetAccountByEmail(ctx context.Context, email string) (*account.Account, error) {
	acc, err := a.accountRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, account.ErrAccountNotFound) {
			return nil, apperror.WithCause(apperror.NotFound, err)
		}
		return nil, apperror.WithCause(apperror.Internal, err)
	}

	return acc, nil
}
