package banner

import (
	"strings"

	"github.com/ming-0x0/yuan/internal/modules/common/domain"
	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
)

type Banner = *banner

type banner struct {
	ID            id.ID
	NameVI        string
	NameEN        string
	NameZH        string
	DescriptionVI string
	DescriptionEN string
	DescriptionZH string
	Position      int
	Status        int
	ResourceID    id.ID
	Link          string
	ButtonNameVI  string
	ButtonNameEN  string
	ButtonNameZH  string
	HasContent    bool
}

func New(nameVI string, resourceID id.ID) (Banner, error) {
	if strings.TrimSpace(nameVI) == "" {
		return nil, domain.ErrRequiredField
	}
	return &banner{
		ID:         id.MustNew(),
		NameVI:     nameVI,
		ResourceID: resourceID,
		Status:     1,
	}, nil
}
