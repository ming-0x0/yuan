package permissiongroup

import (
	"strings"

	"github.com/ming-0x0/yuan/internal/common/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/iam/domain/permission"
)

type PermissionGroup struct {
	ID             id.ID
	Name           string
	Description    string
	FullPermission bool
	Permissions    []*permission.Permission
}

func New(name, description string, fullPermission bool, permissions []*permission.Permission) (*PermissionGroup, error) {
	if strings.TrimSpace(name) == "" {
		return nil, domain.ErrRequiredField
	}
	return &PermissionGroup{
		ID:             id.MustNew(),
		Name:           name,
		Description:    description,
		FullPermission: fullPermission,
		Permissions:    permissions,
	}, nil
}
