package account

import (
	"context"

	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type AccountRepository interface {
	Create(ctx context.Context, account Account) error
	Update(ctx context.Context, account Account) error
	FindByEmail(ctx context.Context, email string) (Account, error)
	FindByID(ctx context.Context, id id.ID) (Account, error)
}

type Account = *account

type account struct {
	ID             id.ID
	Email          string
	HashedPassword string
	FullName       string
}

func New(
	email string,
	hashedPassword string,
	fullName string,
) (Account, error) {
	account := &account{
		ID:             id.MustNew(),
		Email:          email,
		HashedPassword: hashedPassword,
		FullName:       fullName,
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
	fullName string,
) (Account, error) {
	account := &account{
		ID:             id,
		Email:          email,
		HashedPassword: hashedPassword,
		FullName:       fullName,
	}

	if err := account.validate(); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *account) validate() error {
	return nil
}
