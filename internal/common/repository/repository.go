package repository

import (
	"context"
)

type Condition any

//go:generate go tool mockgen -source=repository.go -destination=mocks/repository.go -package=mocks
type TransactionInterface interface {
	Execute(ctx context.Context, fn func(ctx context.Context) error) error
}

type Model interface {
	Create(createdBy int64)
	Update(updatedBy int64)
	Delete(deletedBy int64)
}

type Repository[D any] interface {
	Create(ctx context.Context, domain *D, createdBy int64) error
	TakeByConditions(ctx context.Context, conditions ...Condition) (*D, error)
	FindByConditions(ctx context.Context, conditions ...Condition) ([]*D, error)
	Update(ctx context.Context, domain *D, updatedBy int64) error
	Delete(ctx context.Context, domain *D, deletedBy int64) error
}

type Converter[D, M any] interface {
	ToDomain(*M) (*D, error)
	ToModel(*D) (*M, error)
	ToDomainSlice([]*M) ([]*D, error)
	ToModelSlice([]*D) ([]*M, error)
}
