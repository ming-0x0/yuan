package internal

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64  `bun:"id,pk,type:bigint"`
	Email         string `bun:"email,type:varchar(265),notnull,unique"`
	Username      string `bun:"username,type:varchar(265),notnull,unique"`
	Password      string `bun:"password,type:varchar(100),notnull"`
	BaseModelWithDeletedAt
}
