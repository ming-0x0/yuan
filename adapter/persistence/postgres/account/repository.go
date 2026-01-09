package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ming-0x0/yuan/adapter/persistence/postgres/client"
	"github.com/ming-0x0/yuan/adapter/persistence/postgres/internal"
	"github.com/ming-0x0/yuan/internal/account/domain/account"
	"github.com/ming-0x0/yuan/pkg/logger"
	"github.com/uptrace/bun"
)

type AccountRepository struct {
	client *client.Client
	logger *logger.Logger
}

func New(client *client.Client, logger *logger.Logger) *AccountRepository {
	return &AccountRepository{client: client, logger: logger}
}

func (r *AccountRepository) FindByEmail(ctx context.Context, email string) (*account.Account, error) {
	internalAccount := new(internal.Account)

	err := r.client.DB(ctx).NewSelect().Model(internalAccount).Where("? = ?", bun.Ident("email"), email).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, account.ErrAccountNotFound
		}
		return nil, err
	}

	return ToDomain(internalAccount)
}

func (r *AccountRepository) Create(ctx context.Context, account *account.Account) error {
	internalAccount, err := ToModel(account)
	if err != nil {
		return err
	}

	_, err = r.client.DB(ctx).NewInsert().Model(internalAccount).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
