package account

import (
	"context"

	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/common/validator"
	"github.com/ming-0x0/yuan/internal/common/validator/rule"
)

type AccountRepositoryInterface interface {
	Create(ctx context.Context, account Account) error
	FindByEmail(ctx context.Context, email string) (Account, error)
}

type Account = *account

type account struct {
	ID             id.ID
	Email          string
	HashedPassword string
}

func New(
	email string,
	hashedPassword string,
) (Account, error) {
	account := &account{
		ID:             id.MustNew(),
		Email:          email,
		HashedPassword: hashedPassword,
	}

	if err := account.validate(); err != nil {
		return nil, err
	}

	return account, nil
}

func FromRepository(
	id id.ID,
	email string,
	hashedPassword string,
) (Account, error) {
	account := &account{
		ID:             id,
		Email:          email,
		HashedPassword: hashedPassword,
	}

	if err := account.validate(); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *account) validate() error {
	return validator.New().
		Assert(rule.Required(a.ID)).
		Assert(rule.Required(a.Email)).
		Assert(rule.IsEmail(a.Email)).
		Assert(rule.Required(a.HashedPassword)).
		Err()
}
