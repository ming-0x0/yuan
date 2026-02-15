package application

import (
	"context"
	"errors"

	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain/user"
)

type IAMService interface {
	Register(ctx context.Context, req RegisterRequest) (id.ID, error)
	Login(ctx context.Context, username, password string) (string, error)
	CheckPermission(ctx context.Context, userID id.ID, permCode string) (bool, error)
}

type RegisterRequest struct {
	FullName string
	Email    string
	Username string
	Password string
}

type iamService struct {
	userRepo domain.UserRepository
	pgRepo   domain.PermissionGroupRepository
}

func NewIAMService(userRepo domain.UserRepository, pgRepo domain.PermissionGroupRepository) IAMService {
	return &iamService{
		userRepo: userRepo,
		pgRepo:   pgRepo,
	}
}

func (s *iamService) Register(ctx context.Context, req RegisterRequest) (id.ID, error) {
	u, _ := s.userRepo.FindByUsername(ctx, req.Username)
	if u != nil {
		return 0, errors.New("user already exists")
	}

	newUser, err := user.New(req.FullName, req.Email, req.Username, req.Password, false, nil)
	if err != nil {
		return 0, err
	}

	if err := s.userRepo.Save(ctx, newUser); err != nil {
		return 0, err
	}

	return newUser.ID, nil
}

func (s *iamService) Login(ctx context.Context, username, password string) (string, error) {
	u, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil || u == nil {
		return "", errors.New("invalid credentials")
	}

	if !u.CanLogin() {
		return "", errors.New("account is blocked")
	}

	if u.Password != password {
		return "", errors.New("invalid credentials")
	}

	return "jwt-token-for-" + u.Username, nil
}

func (s *iamService) CheckPermission(ctx context.Context, userID id.ID, permCode string) (bool, error) {
	u, err := s.userRepo.FindByID(ctx, userID)
	if err != nil || u == nil {
		return false, errors.New("user not found")
	}

	return u.HasPermission(permCode), nil
}
