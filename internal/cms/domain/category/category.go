package category

import (
	"strings"

	"github.com/ming-0x0/yuan/internal/common/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type Category = *category

type category struct {
	ID            id.ID
	NameVI        string
	NameEN        string
	NameZH        string
	DescriptionVI string
	DescriptionEN string
	DescriptionZH string
	CategoryType  int
	RouterVI      string
	RouterEN      string
	RouterZH      string
	Position      int
	ParentID      id.ID
	Status        int
	ResourceID    id.ID
	Level         int
}

func New(nameVI string, catType int, level int) (Category, error) {
	if strings.TrimSpace(nameVI) == "" {
		return nil, domain.ErrRequiredField
	}
	return &category{
		ID:           id.MustNew(),
		NameVI:       nameVI,
		CategoryType: catType,
		Status:       1,
		Level:        level,
	}, nil
}
