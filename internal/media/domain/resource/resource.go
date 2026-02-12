package resource

import (
	"strings"

	"github.com/ming-0x0/yuan/internal/common/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type Resource = *resource

type resource struct {
	ID           id.ID
	Name         string
	Description  string
	ResourceType int64
	URL          string
	YoutubeID    string
}

func New(name, url string, resType int64) (Resource, error) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(url) == "" {
		return nil, domain.ErrRequiredField
	}
	return &resource{
		ID:           id.MustNew(),
		Name:         name,
		URL:          url,
		ResourceType: resType,
	}, nil
}
