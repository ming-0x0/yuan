package account

import (
	"github.com/ming-0x0/yuan/adapter/persistence/postgres/internal"
	"github.com/ming-0x0/yuan/internal/account/domain/account"
)

func ToDomain(src *internal.Account) (*account.Account, error) {
	return account.New(src.Email, src.Password)
}

func ToModel(src *account.Account) (*internal.Account, error) {
	return &internal.Account{
		ID:       src.ID().Int64(),
		Email:    src.Email(),
		Password: src.HashedPassword(),
	}, nil
}
