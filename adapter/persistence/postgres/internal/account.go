package internal

import (
	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts,alias:a"`
	ID            int64  `bun:"id,pk,type:bigint"`
	Email         string `bun:"email,type:varchar(265),notnull,unique"`
	Password      string `bun:"password,type:varchar(100),notnull"`
	BaseModelWithDeletedAt
}
