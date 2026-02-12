package partner

import (
	"strings"

	"github.com/ming-0x0/yuan/internal/common/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type Partner = *partner

type partner struct {
	ID            id.ID
	Name          string
	DescriptionVI string
	DescriptionEN string
	DescriptionZH string
	ResourceID    id.ID
	Status        int
	Position      int
	Link          string
}

func New(name string, resourceID id.ID) (Partner, error) {
	if strings.TrimSpace(name) == "" {
		return nil, domain.ErrRequiredField
	}
	return &partner{
		ID:         id.MustNew(),
		Name:       name,
		ResourceID: resourceID,
		Status:     1,
	}, nil
}
