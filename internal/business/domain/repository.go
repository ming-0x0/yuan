package domain

import (
	"context"

	"github.com/ming-0x0/yuan/internal/business/domain/partner"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type PartnerRepository interface {
	FindByID(ctx context.Context, id id.ID) (partner.Partner, error)
	List(ctx context.Context) ([]partner.Partner, error)
	Save(ctx context.Context, p partner.Partner) error
}
