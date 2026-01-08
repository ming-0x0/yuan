package domain

import (
	"context"

	"github.com/ming-0x0/yuan/internal/common/validator"
	"github.com/ming-0x0/yuan/internal/common/validator/rule"
)

type AccountRepository interface {
	Create(ctx context.Context, account *Account) error
	FindByEmail(ctx context.Context, email string) (*Account, error)
}

type Account struct {
	id             int64
	email          string
	hashedPassword string
}

func (a *Account) ID() int64 {
	return a.id
}

func (a *Account) Email() string {
	return a.email
}

func (a *Account) HashedPassword() string {
	return a.hashedPassword
}

func NewAccount(
	id int64,
	email string,
	password string,
) (*Account, error) {
	account := &Account{
		id:             id,
		email:          email,
		hashedPassword: password,
	}
	if err := account.validate(); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) validate() error {
	return validator.New().
		Assert(rule.Required(a.id)).
		Assert(rule.Required(a.email)).Yield("email is required").
		Assert(rule.IsEmail(a.email)).Yield("email is invalid").
		Assert(rule.Required(a.hashedPassword)).
		Err()
}
