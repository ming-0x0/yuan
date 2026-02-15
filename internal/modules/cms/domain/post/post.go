package post

import (
	"strings"
	"time"

	"github.com/ming-0x0/yuan/internal/modules/common/domain"
	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
)

type Post = *post

type post struct {
	ID            id.ID
	TitleVI       string
	TitleEN       string
	TitleZH       string
	SlugVI        string
	SlugEN        string
	SlugZH        string
	AltVI         string
	AltEN         string
	AltZH         string
	DescriptionVI string
	DescriptionEN string
	DescriptionZH string
	InfoVI        string
	InfoEN        string
	InfoZH        string
	ResourceIDs   []id.ID
	ContentVI     string
	ContentEN     string
	ContentZH     string
	Status        int
	PostType      int
	CategoryID    id.ID
	PublicDate    *time.Time
}

func New(
	titleVI, slugVI, contentVI string,
	status, postType int,
	categoryID id.ID,
) (Post, error) {
	if strings.TrimSpace(titleVI) == "" || strings.TrimSpace(slugVI) == "" || strings.TrimSpace(contentVI) == "" {
		return nil, domain.ErrRequiredField
	}

	return &post{
		ID:         id.MustNew(),
		TitleVI:    titleVI,
		SlugVI:     slugVI,
		ContentVI:  contentVI,
		Status:     status,
		PostType:   postType,
		CategoryID: categoryID,
	}, nil
}
