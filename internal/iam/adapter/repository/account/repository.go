package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/iam/adapter/repository/internal"
	"github.com/ming-0x0/yuan/internal/iam/domain/account"
	"github.com/ming-0x0/yuan/pkg/logger"
	"github.com/ming-0x0/yuan/pkg/orm/bun/client"
	"github.com/uptrace/bun"
)

type accountRepository struct {
	client *client.Client
	logger logger.Logger
}

func New(client *client.Client, logger logger.Logger) *accountRepository {
	return &accountRepository{
		client: client,
		logger: logger,
	}
}

func (r *accountRepository) Create(ctx context.Context, account account.Account) error {
	acc := ToInternal(account)

	acc.CreatedBy = account.ID.Int64()
	acc.UpdatedBy = account.ID.Int64()

	_, err := r.client.DB(ctx).NewInsert().Model(acc).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *accountRepository) FindByEmail(ctx context.Context, email string) (account.Account, error) {
	acc := new(internal.Account)

	err := r.client.DB(ctx).NewSelect().Model(acc).Where("? = ?", bun.Ident("email"), email).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, account.ErrAccountNotFound
		}
		return nil, err
	}

	return ToDomain(acc)
}

func (r *accountRepository) FindByID(ctx context.Context, id id.ID) (account.Account, error) {
	acc := new(internal.Account)

	err := r.client.DB(ctx).NewSelect().Model(acc).Where("? = ?", bun.Ident("id"), id.Int64()).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, account.ErrAccountNotFound
		}
		return nil, err
	}

	return ToDomain(acc)
}

func (r *accountRepository) Update(ctx context.Context, account account.Account) error {
	acc := ToInternal(account)
	acc.UpdatedBy = account.ID.Int64()

	_, err := r.client.DB(ctx).NewUpdate().Model(acc).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
