package footer

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/godruoyi/go-snowflake"
)

type Footer = *footer

type footer struct {
	ID        uint64
	NameVi    string
	NameEn    string
	NameZh    string
	ContentVi string
	ContentEn string
	ContentZh string
	Link      string
}

func New(nameVi, nameEn, nameZh string) (Footer, error) {
	f := &footer{
		ID:     snowflake.ID(),
		NameVi: nameVi,
		NameEn: nameEn,
		NameZh: nameZh,
	}

	if err := f.validate(); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *footer) validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.NameVi, validation.Required),
		validation.Field(&f.NameEn, validation.Required),
		validation.Field(&f.NameZh, validation.Required),
	)
}
