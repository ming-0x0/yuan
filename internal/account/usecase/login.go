package usecase

import (
	"context"
	"errors"

	"github.com/ming-0x0/yuan/internal/account/domain/account"
	"github.com/ming-0x0/yuan/internal/common/domain/password"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	ID    int64
	Email string
}

type LoginUseCase struct {
	repo   account.Repository
	pwdSvc password.Password
}

func NewLoginUseCase(
	repo account.Repository,
	pwdSvc password.Password,
) *LoginUseCase {
	return &LoginUseCase{
		repo:   repo,
		pwdSvc: pwdSvc,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	acc, err := uc.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return LoginResponse{}, err
	}
	if acc == nil {
		return LoginResponse{}, ErrInvalidCredentials
	}

	match, err := uc.pwdSvc.Compare(acc.HashedPassword(), req.Password)
	if err != nil {
		return LoginResponse{}, err
	}
	if !match {
		return LoginResponse{}, ErrInvalidCredentials
	}

	return LoginResponse{
		ID:    acc.ID(),
		Email: acc.Email(),
	}, nil
}
