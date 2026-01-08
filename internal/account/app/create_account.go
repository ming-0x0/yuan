package app

import (
	"context"

	"github.com/ming-0x0/yuan/internal/account/domain"
	"github.com/ming-0x0/yuan/internal/common/apperror"
)

func (a *accountApp) CreateAccount(ctx context.Context, email string, password string) error {
	accountID, err := sf.NextID()
	if err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	account, err := domain.NewAccount(accountID, email, password)
	if err != nil {
		return apperror.WithCause(apperror.InvalidArgument, err)
	}

	return a.accountRepo.Create(ctx, account)
}
