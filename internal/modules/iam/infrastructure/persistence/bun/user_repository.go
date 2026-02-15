package bun

import (
	"context"
	"database/sql"

	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain/user"
	orm "github.com/ming-0x0/yuan/pkg/orm/bun/db"
)

type userRepository struct {
	db *orm.DB
}

func NewUserRepository(db *orm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, u *user.User) error {
	m := userFromDomain(u)
	_, err := r.db.WithContext(ctx).NewInsert().Model(m).
		On("DUPLICATE KEY UPDATE").
		Exec(ctx)
	return err
}

func (r *userRepository) FindByID(ctx context.Context, id id.ID) (*user.User, error) {
	m := new(userModel)
	err := r.db.WithContext(ctx).NewSelect().
		Model(m).
		Relation("PermissionGroup").
		Relation("PermissionGroup.Permissions").
		Where("u.id = ?", id).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return userToDomain(m), nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	m := new(userModel)
	err := r.db.WithContext(ctx).NewSelect().
		Model(m).
		Relation("PermissionGroup").
		Relation("PermissionGroup.Permissions").
		Where("u.username = ?", username).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return userToDomain(m), nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	m := new(userModel)
	err := r.db.WithContext(ctx).NewSelect().
		Model(m).
		Relation("PermissionGroup").
		Where("u.email = ?", email).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return userToDomain(m), nil
}
