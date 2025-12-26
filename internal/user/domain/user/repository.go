package user

import "context"

type UserRepositoryInterface interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User, saveBy int64) error
}
