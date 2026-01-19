package auth

import (
	"context"
	"errors"

	"github.com/ming-0x0/yuan/internal/auth/domain/account"
	"github.com/ming-0x0/yuan/internal/common/apperror"
	"golang.org/x/crypto/bcrypt"
)

func (a *AuthApp) Login(ctx context.Context, email string, password string) error {
	acc, err := a.accountRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, account.ErrAccountNotFound) {
			return apperror.WithCause(apperror.Unauthenticated, err)
		}

		return apperror.WithCause(apperror.Internal, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.HashedPassword()), []byte(password))
	if err != nil {
		return apperror.WithCause(apperror.Unauthenticated, err)
	}

	return nil
}
