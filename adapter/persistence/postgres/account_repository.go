package postgres

import (
	"context"
	"database/sql"

	"github.com/ming-0x0/yuan/internal/account/domain/account"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(ctx context.Context, a *account.Account) error {
	query := `INSERT INTO accounts (id, email, hashed_password) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, a.ID(), a.Email(), a.HashedPassword())
	return err
}

func (r *AccountRepository) FindByEmail(ctx context.Context, email string) (*account.Account, error) {
	query := `SELECT id, email, hashed_password FROM accounts WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)

	var id int64
	var foundEmail, hashedPassword string
	if err := row.Scan(&id, &foundEmail, &hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Or a specific error like ErrAccountNotFound
		}
		return nil, err
	}

	return account.New(id, foundEmail, hashedPassword)
}
