package transaction

import "context"

//go:generate go tool mockgen -source=transaction.go -destination=mocks/transaction.go -package=mocks
type TransactionInterface interface {
	Execute(ctx context.Context, fn func(ctx context.Context) error) error
}
