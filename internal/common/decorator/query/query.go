package query

import (
	"context"
	"fmt"
	"strings"

	"github.com/ming-0x0/yuan/pkg/logger"
)

func Execute[Q any, R any](query Query[Q, R], logger logger.Logger) Query[Q, R] {
	return defaultQueryLoggingDecorator[Q, R]{
		query:  query,
		logger: logger,
	}
}

type Query[Q any, R any] interface {
	Handle(ctx context.Context, query Q) (R, error)
}

type defaultQueryLoggingDecorator[Q any, R any] struct {
	query  Query[Q, R]
	logger logger.Logger
}

func (d defaultQueryLoggingDecorator[Q, R]) Handle(ctx context.Context, query Q) (result R, err error) {
	logger := d.logger.WithFields(
		"query", generateQueryName(query),
		"query_body", fmt.Sprintf("%#v", query),
	)

	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("Query executed successfully")
		} else {
			logger.Error("Failed to execute query", "error", err)
		}
	}()

	return d.query.Handle(ctx, query)
}

func generateQueryName(query any) string {
	return strings.Split(fmt.Sprintf("%T", query), ".")[1]
}
