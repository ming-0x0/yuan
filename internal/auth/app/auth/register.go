package auth

import (
	"context"

	"github.com/ming-0x0/yuan/internal/auth/domain/account"
	"github.com/ming-0x0/yuan/internal/common/apperror"
	"golang.org/x/crypto/bcrypt"
)

func (a *AuthApp) Register(ctx context.Context, email string, password string) error {
	acc, err := a.accountRepo.FindByEmail(ctx, email)
	if err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	if acc != nil {
		return apperror.WithCause(apperror.AlreadyExists, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	account, err := account.New(email, string(hashedPassword))
	if err != nil {
		return apperror.WithCause(apperror.InvalidArgument, err)
	}

	err = a.accountRepo.Create(ctx, account)
	if err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	return nil
}
