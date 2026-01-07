package usecase

import (
	"context"

	"github.com/ming-0x0/yuan/internal/account/domain/account"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/common/domain/password"
)

type RegisterRequest struct {
	Email    string
	Password string
}

type RegisterResponse struct {
	ID    int64
	Email string
}

type RegisterUseCase struct {
	repo     account.Repository
	pwdSvc   password.Password
	idGen    id.ID
}

func NewRegisterUseCase(
	repo account.Repository,
	pwdSvc password.Password,
	idGen id.ID,
) *RegisterUseCase {
	return &RegisterUseCase{
		repo:   repo,
		pwdSvc: pwdSvc,
		idGen:  idGen,
	}
}

func (uc *RegisterUseCase) Execute(ctx context.Context, req RegisterRequest) (RegisterResponse, error) {
	hashedPassword, err := uc.pwdSvc.Hash(req.Password)
	if err != nil {
		return RegisterResponse{}, err
	}

	nextID, err := uc.idGen.Next()
	if err != nil {
		return RegisterResponse{}, err
	}

	acc, err := account.New(nextID, req.Email, hashedPassword)
	if err != nil {
		return RegisterResponse{}, err
	}

	if err := uc.repo.Save(ctx, acc); err != nil {
		return RegisterResponse{}, err
	}

	return RegisterResponse{
		ID:    acc.ID(),
		Email: acc.Email(),
	}, nil
}
