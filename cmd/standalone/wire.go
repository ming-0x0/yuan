//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/ming-0x0/yuan/adapter/external/crypto/bcrypt"
	accHttp "github.com/ming-0x0/yuan/adapter/delivery/http/account"
	"github.com/ming-0x0/yuan/adapter/external/id/sonyflake"
	postgresRepo "github.com/ming-0x0/yuan/adapter/persistence/postgres"
	"github.com/ming-0x0/yuan/internal/account/domain/account"
	"github.com/ming-0x0/yuan/internal/account/usecase"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/common/domain/password"
)

func ProvideBcrypt() *bcrypt.Bcrypt {
	return bcrypt.New(10)
}

func InitializeApp(db *sql.DB) (*accHttp.Handler, error) {
	wire.Build(
		// Adapters
		sonyflake.New,
		ProvideBcrypt,
		postgresRepo.NewAccountRepository,

		// Bind interfaces
		wire.Bind(new(account.Repository), new(*postgresRepo.AccountRepository)),
		wire.Bind(new(password.Password), new(*bcrypt.Bcrypt)),
		wire.Bind(new(id.ID), new(*sonyflake.Sonyflake)),

		// Use cases
		usecase.NewRegisterUseCase,
		usecase.NewLoginUseCase,

		// Delivery
		accHttp.NewHandler,
	)
	return nil, nil
}
