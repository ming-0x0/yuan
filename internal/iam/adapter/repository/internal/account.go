package internal

import (
	"github.com/ming-0x0/yuan/pkg/orm/bun/model"
	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts,alias:a"`
	ID            int64  `bun:"id,pk,type:bigint"`
	Email         string `bun:"email,notnull,type:varchar(265),unique"`
	Password      string `bun:"password,notnull,type:varchar(255)"`
	FullName      string `bun:"full_name,type:varchar(255)"`
	model.MetadataWithDeletedAt
}
