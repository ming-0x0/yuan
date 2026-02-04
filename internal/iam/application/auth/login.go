package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ming-0x0/yuan/internal/iam/domain/account"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Login(ctx context.Context, email string, password string) (string, error) {
	acc, err := s.accountRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if acc == nil {
		return "", account.ErrAccountNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acc.HashedPassword), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	// In a real app, strict signing key management is required.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": acc.ID.String(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	// TODO: Move secret to config
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
