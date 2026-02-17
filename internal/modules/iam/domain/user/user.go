package user

import (
	"net/mail"
	"strings"

	"github.com/ming-0x0/yuan/internal/modules/common/domain"
	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain/permissiongroup"
)

type Status uint8

const (
	Activate Status = 1
	Inactive Status = 2
)

type User = *user

type user struct {
	ID              id.ID
	FullName        string
	Email           string
	Username        string
	Password        string
	IsAdmin         bool
	Status          Status
	ReceiveEmail    bool
	PermissionGroup *permissiongroup.PermissionGroup
}

func New(
	fullName, email, username, password string,
	isAdmin bool,
	pg *permissiongroup.PermissionGroup,
) (User, error) {
	if strings.TrimSpace(fullName) == "" {
		return nil, domain.ErrRequiredField
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, domain.ErrInvalidEmail
	}
	if strings.TrimSpace(username) == "" {
		return nil, domain.ErrRequiredField
	}
	if len(password) < 8 {
		return nil, domain.ErrInvalidFormat
	}

	return &user{
		ID:              id.MustNew(),
		FullName:        fullName,
		Email:           email,
		Username:        username,
		Password:        password,
		IsAdmin:         isAdmin,
		Status:          Activate,
		ReceiveEmail:    true,
		PermissionGroup: pg,
	}, nil
}

func (u User) HasPermission(code string) bool {
	if u.IsAdmin {
		return true
	}
	if u.PermissionGroup == nil {
		return false
	}
	if u.PermissionGroup.FullPermission {
		return true
	}
	for _, p := range u.PermissionGroup.Permissions {
		if p.Code == code {
			return true
		}
	}
	return false
}

func (u User) Activate() {
	u.Status = Activate
}

func (u User) Block() {
	u.Status = Inactive
}

func (u User) CanLogin() bool {
	return u.Status == Activate
}
