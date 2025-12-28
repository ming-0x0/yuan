package client

import (
	"context"
	"database/sql"

	ctxkey "github.com/ming-0x0/yuan/pkg/ctxutil/key"
	"github.com/ming-0x0/yuan/pkg/logger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	_ "github.com/lib/pq"
)

type Client struct {
	db     *bun.DB
	logger logger.Logger
}

func New(
	db *sql.DB,
	logger logger.Logger,
) (*Client, error) {
	return &Client{
		db:     bun.NewDB(db, pgdialect.New()),
		logger: logger,
	}, nil
}

func (c *Client) DB(ctx context.Context) bun.IDB {
	v := ctx.Value(ctxkey.TransactionContextKey)
	if v != nil {
		if tx, ok := v.(bun.Tx); ok {
			return tx
		}
	}
	return c.db
}
