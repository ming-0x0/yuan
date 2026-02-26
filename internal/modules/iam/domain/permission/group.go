package permission

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/godruoyi/go-snowflake"
)

type PermissionGroup = *permissionGroup

type permissionGroup struct {
	ID             uint64
	Name           string
	Description    string
	FullPermission bool
}

func NewGroup(name, description string, fullPermission bool) (PermissionGroup, error) {
	pg := &permissionGroup{
		ID:             snowflake.ID(),
		Name:           name,
		Description:    description,
		FullPermission: fullPermission,
	}

	if err := pg.validate(); err != nil {
		return nil, err
	}

	return pg, nil
}

func (pg *permissionGroup) validate() error {
	return validation.ValidateStruct(pg,
		validation.Field(&pg.Name, validation.Required),
	)
}
