package account

import (
	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/iam/adapter/repository/internal"
	"github.com/ming-0x0/yuan/internal/iam/domain/account"
)

func ToDomain(src *internal.Account) (account.Account, error) {
	return account.FromRepository(
		id.FromInt64(src.ID),
		src.Email,
		src.Password,
	)
}

func ToInternal(src account.Account) *internal.Account {
	return &internal.Account{
		ID:       src.ID.Int64(),
		Email:    src.Email,
		Password: src.HashedPassword,
	}
}
