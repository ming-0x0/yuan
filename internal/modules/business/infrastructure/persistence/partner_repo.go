package persistence

import (
	"context"

	"github.com/ming-0x0/yuan/internal/modules/business/domain"
	"github.com/ming-0x0/yuan/internal/modules/business/domain/partner"
	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
)

type partnerRepository struct{}

func NewPartnerRepository() domain.PartnerRepository {
	return &partnerRepository{}
}

func (r *partnerRepository) FindByID(ctx context.Context, id id.ID) (partner.Partner, error) {
	return nil, nil // TODO: Implement
}

func (r *partnerRepository) List(ctx context.Context) ([]partner.Partner, error) {
	return nil, nil // TODO: Implement
}

func (r *partnerRepository) Save(ctx context.Context, p partner.Partner) error {
	return nil // TODO: Implement
}
