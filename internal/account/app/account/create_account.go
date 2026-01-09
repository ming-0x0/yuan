package app

import (
	"context"

	"github.com/ming-0x0/yuan/internal/account/domain/account"
	"github.com/ming-0x0/yuan/internal/common/apperror"
	"golang.org/x/crypto/bcrypt"
)

func (a *AccountApp) CreateAccount(ctx context.Context, email string, password string) error {
	accountID, err := sf.NextID()
	if err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	account, err := account.New(accountID, email, string(hashedPassword))
	if err != nil {
		return apperror.WithCause(apperror.InvalidArgument, err)
	}

	err = a.accountRepo.Create(ctx, account)
	if err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	return nil
}
