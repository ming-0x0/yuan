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

func (m *BaseModel) Create(createdBy int64) {
	m.CreatedBy = createdBy
	m.UpdatedBy = createdBy
}

func (m *BaseModel) Update(updatedBy int64) {
	m.UpdatedBy = updatedBy
	m.UpdatedAt = time.Now()
}

func (m *BaseModelWithDeletedAt) Delete(deletedBy int64) {
	m.DeletedBy = void.New(deletedBy)
	m.DeletedAt = void.New(time.Now())
}
