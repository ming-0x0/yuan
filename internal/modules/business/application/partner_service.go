package application

import (
	"context"

	"github.com/ming-0x0/yuan/internal/modules/business/domain"
	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
)

type PartnerDTO struct {
	ID         id.ID  `json:"id"`
	Name       string `json:"name"`
	ResourceID id.ID  `json:"resource_id"`
}

type PartnerService interface {
	GetPartner(ctx context.Context, id id.ID) (*PartnerDTO, error)
	ListPartners(ctx context.Context) ([]*PartnerDTO, error)
}

type partnerService struct {
	repo domain.PartnerRepository
}

func NewPartnerService(repo domain.PartnerRepository) PartnerService {
	return &partnerService{
		repo: repo,
	}
}

func (s *partnerService) GetPartner(ctx context.Context, id id.ID) (*PartnerDTO, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, nil
	}

	return &PartnerDTO{
		ID:         p.ID,
		Name:       p.Name,
		ResourceID: p.ResourceID,
	}, nil
}

func (s *partnerService) ListPartners(ctx context.Context) ([]*PartnerDTO, error) {
	partners, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var dtos []*PartnerDTO
	for _, p := range partners {
		dtos = append(dtos, &PartnerDTO{
			ID:         p.ID,
			Name:       p.Name,
			ResourceID: p.ResourceID,
		})
	}

	return dtos, nil
}
