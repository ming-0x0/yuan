package permission

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/godruoyi/go-snowflake"
)

type Permission = *permission

type permission struct {
	ID             uint64
	PermissionName string
	FunctionCode   string
}

func New(permissionName, functionCode string) (Permission, error) {
	p := &permission{
		ID:             snowflake.ID(),
		PermissionName: permissionName,
		FunctionCode:   functionCode,
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *permission) validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.PermissionName, validation.Required),
		validation.Field(&p.FunctionCode, validation.Required),
	)
}
