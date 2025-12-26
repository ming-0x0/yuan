package client

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"time"

	"github.com/ming-0x0/yuan/config"
	"github.com/ming-0x0/yuan/pkg/logger"
	bunlogger "github.com/ming-0x0/yuan/pkg/logger/bun"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Client struct {
	db     *bun.DB
	logger logger.Logger
}

func New(
	config *config.PostgresConfig,
	bunConfig *config.BunConfig,
	db *sql.DB,
	logger logger.Logger,
) *Client {
	pgConfig := &pgdriver.Config{
		Network: "tcp",
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial(network, addr)
		},
		User:       config.User,
		Password:   config.Password,
		Database:   config.Database,
		BufferSize: 1024 * 1024, // 1MB to avoid bufio.Scanner: token too long
	}

	bunDB := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithConfig(pgConfig))), pgdialect.New())

	bunDB.AddQueryHook(bunlogger.New(
		bunlogger.WithDriver("logrus"),
		bunlogger.WithLogger(logger),
		bunlogger.WithSlowThreshold(time.Duration(bunConfig.SlowThreshold)*time.Millisecond),
		bunlogger.WithIgnoreNoRowsError(),
		bunlogger.WithLevel(bunConfig.Level),
	))

	return &Client{
		db:     bunDB,
		logger: logger,
	}
}
