package token

import (
	"time"

	"github.com/ming-0x0/yuan/internal/common/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type Token = *token

type token struct {
	ID        id.ID
	UserID    id.ID
	Value     string
	ExpiredAt time.Time
}

func New(userID id.ID, value string, duration time.Duration) (Token, error) {
	if value == "" {
		return nil, domain.ErrRequiredField
	}
	return &token{
		ID:        id.MustNew(),
		UserID:    userID,
		Value:     value,
		ExpiredAt: time.Now().Add(duration),
	}, nil
}
