package banner

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/godruoyi/go-snowflake"
)

type Status uint8

const (
	StatusActive   Status = 1
	StatusInactive Status = 2
)

type Banner = *banner

type banner struct {
	ID            uint64
	NameVi        string
	NameEn        string
	NameZh        string
	DescriptionVi string
	DescriptionEn string
	DescriptionZh string
	Position      int
	Status        Status
	ResourceID    uint64
	Link          string
	ButtonNameVi  string
	ButtonNameEn  string
	ButtonNameZh  string
	HasContent    bool
}

func New(
	nameVi, nameEn, nameZh string,
	resourceID uint64,
	status Status,
) (Banner, error) {
	b := &banner{
		ID:         snowflake.ID(),
		NameVi:     nameVi,
		NameEn:     nameEn,
		NameZh:     nameZh,
		ResourceID: resourceID,
		Status:     status,
	}

	if err := b.validate(); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *banner) validate() error {
	return validation.ValidateStruct(b,
		validation.Field(&b.NameVi, validation.Required),
		validation.Field(&b.NameEn, validation.Required),
		validation.Field(&b.NameZh, validation.Required),
		validation.Field(&b.ResourceID, validation.Required),
		validation.Field(&b.Status, validation.Required),
	)
}
