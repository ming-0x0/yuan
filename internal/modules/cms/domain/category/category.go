package category

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/godruoyi/go-snowflake"
)

type Status uint8

const (
	StatusActive   Status = 1
	StatusInactive Status = 2
)

type Category = *category

type category struct {
	ID            uint64
	NameVi        string
	NameEn        string
	NameZh        string
	DescriptionVi string
	DescriptionEn string
	DescriptionZh string
	Type          int
	RouterVi      string
	RouterEn      string
	RouterZh      string
	Position      int
	ParentID      *uint64
	Status        Status
	ResourceID    *uint64
	Level         int
}

func New(
	nameVi, nameEn, nameZh string,
	categoryType int,
	level int,
	status Status,
) (Category, error) {
	c := &category{
		ID:     snowflake.ID(),
		NameVi: nameVi,
		NameEn: nameEn,
		NameZh: nameZh,
		Type:   categoryType,
		Level:  level,
		Status: status,
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *category) validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.NameVi, validation.Required),
		validation.Field(&c.NameEn, validation.Required),
		validation.Field(&c.NameZh, validation.Required),
		validation.Field(&c.Type, validation.Required),
		validation.Field(&c.Level, validation.Min(0)),
	)
}
