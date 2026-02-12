package footer

import (
	"strings"

	"github.com/ming-0x0/yuan/internal/common/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type Footer = *footer

type footer struct {
	ID        id.ID
	NameVI    string
	NameEN    string
	NameZH    string
	ContentVI string
	ContentEN string
	ContentZH string
	Link      string
}

func New(nameVI string) (Footer, error) {
	if strings.TrimSpace(nameVI) == "" {
		return nil, domain.ErrRequiredField
	}
	return &footer{
		ID:     id.MustNew(),
		NameVI: nameVI,
	}, nil
}
