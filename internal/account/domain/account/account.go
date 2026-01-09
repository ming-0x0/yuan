package account

import (
	"context"
	"errors"

	"github.com/ming-0x0/yuan/internal/common/validator"
	"github.com/ming-0x0/yuan/internal/common/validator/rule"
	"golang.org/x/crypto/bcrypt"
)

type AccountApp interface {
	CreateAccount(ctx context.Context, email string, password string) error
	Login(ctx context.Context, email string, password string) (*Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*Account, error)
}

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

func (a *Account) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.hashedPassword), []byte(password))
}

func New(
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

var (
	ErrAccountNotFound = errors.New("account not found")
)

func (a *Account) validate() error {
	return validator.New().
		Assert(rule.Required(a.id)).
		Assert(rule.Required(a.email)).Yield("email is required").
		Assert(rule.IsEmail(a.email)).Yield("email is invalid").
		Assert(rule.Required(a.hashedPassword)).
		Err()
}
