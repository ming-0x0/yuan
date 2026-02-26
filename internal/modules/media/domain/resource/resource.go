package resource

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/godruoyi/go-snowflake"
)

type Resource = *resource

type resource struct {
	ID          uint64
	Name        string
	Description string
	Type        int
	URL         string
	YoutubeID   string
}

func New(name, url string, resourceType int) (Resource, error) {
	r := &resource{
		ID:   snowflake.ID(),
		Name: name,
		Type: resourceType,
		URL:  url,
	}

	if err := r.validate(); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *resource) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.URL, validation.Required),
		validation.Field(&r.Type, validation.Required),
	)
}
