package account

import (
	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/common/validator"
	"github.com/ming-0x0/yuan/internal/common/validator/rule"
)

type Account struct {
	id             id.ID
	email          string
	hashedPassword string
}

func New(
	id id.ID,
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

func (a *Account) ID() id.ID {
	return a.id
}

func (a *Account) Email() string {
	return a.email
}

func (a *Account) HashedPassword() string {
	return a.hashedPassword
}

func (a *Account) validate() error {
	return validator.New().
		Assert(rule.Required(a.id)).
		Assert(rule.Required(a.email)).
		Assert(rule.IsEmail(a.email)).
		Assert(rule.Required(a.hashedPassword)).
		Err()
}
