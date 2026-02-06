package user

import (
	"context"

	"github.com/ming-0x0/yuan/internal/iam/domain/account"
)

type UserService interface {
	UpdateProfile(ctx context.Context, fullName string) error
	GetProfile(ctx context.Context) (account.Account, error)
}

type userService struct {
	accountRepo account.AccountRepository
}

func NewUserService(accountRepo account.AccountRepository) UserService {
	return &userService{
		accountRepo: accountRepo,
	}
}

func (s *userService) UpdateProfile(ctx context.Context, fullName string) error {
	// TODO: Get userID from context
	// For now, we'll need to implement context extraction
	return nil
}

func (s *userService) GetProfile(ctx context.Context) (account.Account, error) {
	// TODO: Get userID from context
	// For now, we'll need to implement context extraction
	return nil, nil
}
