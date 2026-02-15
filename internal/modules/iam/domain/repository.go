package domain

import (
	"context"

	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain/permissiongroup"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain/user"
)

type UserRepository interface {
	Save(ctx context.Context, u *user.User) error
	FindByID(ctx context.Context, id id.ID) (*user.User, error)
	FindByUsername(ctx context.Context, username string) (*user.User, error)
	FindByEmail(ctx context.Context, email string) (*user.User, error)
}

type PermissionGroupRepository interface {
	FindByID(ctx context.Context, id id.ID) (*permissiongroup.PermissionGroup, error)
	ListAll(ctx context.Context) ([]*permissiongroup.PermissionGroup, error)
}
