package model

import (
	"database/sql"
	"time"
)

type Metadata struct {
	CreatedBy int64     `bun:"created_by,notnull,type:bigint"`
	CreatedAt time.Time `bun:"created_at,notnull,type:timestamp,default:current_timestamp"`
	UpdatedBy int64     `bun:"updated_by,notnull,type:bigint"`
	UpdatedAt time.Time `bun:"updated_at,notnull,type:timestamp,default:current_timestamp"`
}

type MetadataWithDeletedAt struct {
	Metadata
	DeletedBy sql.NullInt64 `bun:"deleted_by,nullzero,type:bigint"`
	DeletedAt sql.NullTime  `bun:"deleted_at,soft_delete,nullzero,type:timestamp"`
}
