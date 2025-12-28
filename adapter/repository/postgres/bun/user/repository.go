package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ming-0x0/yuan/adapter/repository/postgres/bun/client"
	"github.com/ming-0x0/yuan/adapter/repository/postgres/bun/internal"
	"github.com/ming-0x0/yuan/internal/common/apperror"
	"github.com/ming-0x0/yuan/internal/user/domain/user"
	userDTO "github.com/ming-0x0/yuan/internal/user/dto/user"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	client *client.Client
}

func New(client *client.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (userDTO.User, error) {
	user := new(internal.User)
	if err := r.client.DB(ctx).NewSelect().Model(user).Where("? = ?", bun.Ident("id"), id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userDTO.User{}, apperror.WithCause(apperror.NotFound, err)
		}
		return userDTO.User{}, apperror.WithCause(apperror.Internal, err)
	}

	return ToDTO(user)
}

func (r *UserRepository) Create(ctx context.Context, user *user.User, createdBy int64) error {
	model, err := ToModel(user)
	if err != nil {
		return err
	}

	model.CreatedBy = createdBy
	model.UpdatedBy = createdBy

	if _, err := r.client.DB(ctx).NewInsert().Model(model).Exec(ctx); err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	return nil
}
