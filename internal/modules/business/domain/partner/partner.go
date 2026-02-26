package partner

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/godruoyi/go-snowflake"
)

type Status uint8

const (
	StatusActive   Status = 1
	StatusInactive Status = 2
)

type Partner = *partner

type partner struct {
	ID            uint64
	Name          string
	DescriptionVi string
	DescriptionEn string
	DescriptionZh string
	ResourceID    uint64
	Status        Status
	Position      int
	Link          string
}

func New(name string, resourceID uint64, status Status) (Partner, error) {
	p := &partner{
		ID:         snowflake.ID(),
		Name:       name,
		ResourceID: resourceID,
		Status:     status,
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *partner) validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.ResourceID, validation.Required),
		validation.Field(&p.Status, validation.Required),
	)
}
