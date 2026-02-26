package role

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/godruoyi/go-snowflake"
)

type Role = *role

type role struct {
	ID                uint64
	PermissionID      uint64
	PermissionGroupID uint64
}

func New(permissionID, permissionGroupID uint64) (Role, error) {
	r := &role{
		ID:                snowflake.ID(),
		PermissionID:      permissionID,
		PermissionGroupID: permissionGroupID,
	}

	if err := r.validate(); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *role) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.PermissionID, validation.Required),
		validation.Field(&r.PermissionGroupID, validation.Required),
	)
}
