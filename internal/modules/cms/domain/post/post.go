package post

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/godruoyi/go-snowflake"
)

type Status uint8

const (
	StatusActive   Status = 1
	StatusInactive Status = 2
)

type Post = *post

type post struct {
	ID            uint64
	TitleVi       string
	TitleEn       string
	TitleZh       string
	SlugVi        string
	SlugEn        string
	SlugZh        string
	AltVi         string
	AltEn         string
	AltZh         string
	DescriptionVi string
	DescriptionEn string
	DescriptionZh string
	InfoVi        string
	InfoEn        string
	InfoZh        string
	ResourceIDs   string // Could be slice later
	ContentVi     string
	ContentEn     string
	ContentZh     string
	Status        Status
	Type          int
	CategoryID    uint64
	PublicDate    *time.Time
}

func New(
	titleVi, titleEn, titleZh string,
	slugVi, slugEn, slugZh string,
	categoryID uint64,
	status Status,
) (Post, error) {
	p := &post{
		ID:         snowflake.ID(),
		TitleVi:    titleVi,
		TitleEn:    titleEn,
		TitleZh:    titleZh,
		SlugVi:     slugVi,
		SlugEn:     slugEn,
		SlugZh:     slugZh,
		CategoryID: categoryID,
		Status:     status,
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *post) validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.TitleVi, validation.Required),
		validation.Field(&p.SlugVi, validation.Required),
		validation.Field(&p.CategoryID, validation.Required),
	)
}
