package postgres

import (
	"context"
	"database/sql"
	"path/filepath"
	"runtime"
	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresTestContainer struct {
	container *postgres.PostgresContainer
	db        *sql.DB
	connStr   string
}

func (c *PostgresTestContainer) Container() *postgres.PostgresContainer {
	return c.container
}

func (c *PostgresTestContainer) DB() *sql.DB {
	return c.db
}

func (c *PostgresTestContainer) ConnectionString() string {
	return c.connStr
}

func New(ctx context.Context) (*PostgresTestContainer, error) {
	pgContainer, err := postgres.Run(ctx,
		"postgres:18-alpine",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpassword"),
		testcontainers.WithWaitStrategy(
			wait.ForListeningPort("5432/tcp").
				WithStartupTimeout(30*time.Second),
		),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &PostgresTestContainer{
		container: pgContainer,
		db:        db,
		connStr:   connStr,
	}, nil
}

func (c *PostgresTestContainer) Terminate(ctx context.Context) error {
	if c.container != nil {
		return c.container.Terminate(ctx)
	}
	return nil
}

func (c *PostgresTestContainer) Migrate(ctx context.Context) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	goose.SetBaseFS(nil)

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	migrationsDir := filepath.Join(dir, "../../../infrastructure/database/postgres/migrations")

	if err := goose.UpContext(ctx, c.db, migrationsDir); err != nil {
		return err
	}

	return nil
}
