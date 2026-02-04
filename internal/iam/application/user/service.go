package user

import (
	"context"

	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/iam/domain/account"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	UpdateProfile(ctx context.Context, userID id.ID, email string, password string) error
	GetProfile(ctx context.Context, userID id.ID) (account.Account, error)
}

type userService struct {
	accountRepo account.AccountRepository
}

func NewUserService(accountRepo account.AccountRepository) UserService {
	return &userService{
		accountRepo: accountRepo,
	}
}

func (s *userService) UpdateProfile(ctx context.Context, userID id.ID, email string, password string) error {
	acc, err := s.accountRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if email != "" && email != acc.Email {
		acc.Email = email
	}

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		acc.HashedPassword = string(hashedPassword)
	}

	return s.accountRepo.Update(ctx, acc)
}

func (s *userService) GetProfile(ctx context.Context, userID id.ID) (account.Account, error) {
	return s.accountRepo.FindByID(ctx, userID)
}
