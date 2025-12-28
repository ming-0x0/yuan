package user

import "context"

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *User, createdBy int64) error
}
