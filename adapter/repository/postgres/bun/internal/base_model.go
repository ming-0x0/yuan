package internal

import (
	"time"

	"github.com/ming-0x0/yuan/pkg/typeutil/void"
)

type BaseModel struct {
	CreatedBy int64     `bun:"created_by,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedBy int64     `bun:"updated_by,notnull"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

type BaseModelWithDeletedAt struct {
	BaseModel
	DeletedBy void.Void[int64]     `bun:"deleted_by,nullzero,type:bigint"`
	DeletedAt void.Void[time.Time] `bun:"deleted_at,soft_delete,nullzero,type:timestamp"`
}
