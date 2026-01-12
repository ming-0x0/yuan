package app

import (
	"context"
	"errors"

	"github.com/ming-0x0/yuan/internal/account/domain/account"
	"github.com/ming-0x0/yuan/internal/common/apperror"
	"golang.org/x/crypto/bcrypt"
)

func (a *AccountApp) Login(ctx context.Context, email string, password string) (*account.Account, error) {
	acc, err := a.accountRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, account.ErrAccountNotFound) {
			return nil, apperror.WithCause(apperror.Unauthenticated, err)
		}
		return nil, apperror.WithCause(apperror.Internal, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acc.HashedPassword()), []byte(password)); err != nil {
		return nil, apperror.WithCause(apperror.Unauthenticated, err)
	}

	return acc, nil
}
