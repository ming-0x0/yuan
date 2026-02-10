package model

import (
	"database/sql"
	"time"
)

type Timestamp struct {
	CreatedAt time.Time `bun:"created_at,notnull,type:timestamp,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,type:timestamp,default:current_timestamp"`
}

type TimestampWithDeletedAt struct {
	Timestamp
	DeletedAt sql.NullTime `bun:"deleted_at,soft_delete,nullzero,type:timestamp"`
}
