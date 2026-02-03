package client

import (
	"context"
	"database/sql"

	ctxkey "github.com/ming-0x0/yuan/pkg/ctxutil/key"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type Client struct {
	db *bun.DB
}

func NewPostgresClient(sqlDB *sql.DB) (*Client, error) {
	db := bun.NewDB(sqlDB, pgdialect.New())
	return &Client{
		db: db,
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

