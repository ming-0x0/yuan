package account

import (
	"github.com/ming-0x0/yuan/internal/common/validator"
	"github.com/ming-0x0/yuan/internal/common/validator/rule"
)

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

func New(
	id int64,
	email string,
	hashedPassword string,
) (*Account, error) {
	account := &Account{
		id:             id,
		email:          email,
		hashedPassword: hashedPassword,
	}
	if err := account.validate(); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) validate() error {
	return validator.New().
		Assert(rule.Required(a.id)).Yield("id is required").
		Assert(rule.Required(a.email)).Yield("email is required").
		Assert(rule.IsEmail(a.email)).Yield("email is invalid").
		Assert(rule.Required(a.hashedPassword)).Yield("hashedPassword is required").
		Err()
}
