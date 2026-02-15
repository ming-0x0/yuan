package bun

import (
	"time"

	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain/permission"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain/permissiongroup"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain/user"
	"github.com/uptrace/bun"
)

// Bun Models - Shared within the persistence package

type userModel struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID                id.ID     `bun:"id,pk"`
	Email             string    `bun:"email,notnull"`
	Username          string    `bun:"username,notnull"`
	FullName          string    `bun:"full_name"`
	Password          string    `bun:"password"`
	Status            int       `bun:"status"`
	IsAdmin           bool      `bun:"is_admin"`
	PermissionGroupID id.ID     `bun:"permission_group_id"`
	PermissionGroup   *pgModel  `bun:"rel:belongs-to,join:permission_group_id=id"`
	CreatedAt         time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt         time.Time `bun:"updated_at,default:current_timestamp"`
}

type pgModel struct {
	bun.BaseModel `bun:"table:permission_groups,alias:pg"`

	ID             id.ID             `bun:"id,pk"`
	Name           string            `bun:"name,notnull"`
	Description    string            `bun:"description"`
	FullPermission bool              `bun:"full_permission,default:false"`
	Permissions    []permissionModel `bun:"m2m:roles,join:pg=permission"`
}

type permissionModel struct {
	bun.BaseModel `bun:"table:permissions,alias:p"`

	ID   id.ID  `bun:"id,pk"`
	Name string `bun:"name"`
	Code string `bun:"code"`
}

// User Mappings

func userToDomain(m *userModel) *user.User {
	var pg *permissiongroup.PermissionGroup
	if m.PermissionGroup != nil {
		pg = pgToDomain(m.PermissionGroup)
	}

	return &user.User{
		ID:              m.ID,
		FullName:        m.FullName,
		Email:           m.Email,
		Username:        m.Username,
		Password:        m.Password,
		IsAdmin:         m.IsAdmin,
		Status:          user.Status(m.Status),
		PermissionGroup: pg,
	}
}

func userFromDomain(u *user.User) *userModel {
	pgID := id.ID(0)
	if u.PermissionGroup != nil {
		pgID = u.PermissionGroup.ID
	}
	return &userModel{
		ID:                u.ID,
		Email:             u.Email,
		Username:          u.Username,
		FullName:          u.FullName,
		Password:          u.Password,
		Status:            int(u.Status),
		IsAdmin:           u.IsAdmin,
		PermissionGroupID: pgID,
	}
}

// PermissionGroup Mappings

func pgToDomain(m *pgModel) *permissiongroup.PermissionGroup {
	var permissions []*permission.Permission
	for _, p := range m.Permissions {
		permissions = append(permissions, &permission.Permission{
			ID:   p.ID,
			Name: p.Name,
			Code: p.Code,
		})
	}
	return &permissiongroup.PermissionGroup{
		ID:             m.ID,
		Name:           m.Name,
		Description:    m.Description,
		FullPermission: m.FullPermission,
		Permissions:    permissions,
	}
}
