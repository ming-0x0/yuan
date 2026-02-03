package auth

import (
	"context"
	"errors"

	"github.com/ming-0x0/yuan/internal/iam/domain/account"
)

func (s *authService) Register(ctx context.Context, email string, password string) error {
	acc, err := s.accountRepo.FindByEmail(ctx, email)
	if err != nil && !errors.Is(err, account.ErrAccountNotFound) {
		return err
	}

	if acc != nil {
		return errors.New("account already exists")
	}

	acc, err = account.New(email, password)
	if err != nil {
		return err
	}

	err = s.accountRepo.Create(ctx, acc)
	if err != nil {
		return err
	}

	return nil
}
