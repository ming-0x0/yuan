package permission

import (
	"strings"

	"github.com/ming-0x0/yuan/internal/modules/common/domain"
	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
)

type Permission struct {
	ID   id.ID
	Name string
	Code string
}

func New(name, code string) (*Permission, error) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(code) == "" {
		return nil, domain.ErrRequiredField
	}
	return &Permission{
		ID:   id.MustNew(),
		Name: name,
		Code: code,
	}, nil
}
