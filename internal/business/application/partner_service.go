package application

import (
	"context"

	"github.com/ming-0x0/yuan/internal/business/application/port"
	"github.com/ming-0x0/yuan/internal/business/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type PartnerDTO struct {
	ID       id.ID  `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type PartnerService interface {
	GetPartner(ctx context.Context, id id.ID) (*PartnerDTO, error)
	ListPartners(ctx context.Context) ([]*PartnerDTO, error)
}

type partnerService struct {
	repo          domain.PartnerRepository
	mediaProvider port.MediaProvider
}

func NewPartnerService(repo domain.PartnerRepository, mediaProvider port.MediaProvider) PartnerService {
	return &partnerService{
		repo:          repo,
		mediaProvider: mediaProvider,
	}
}

func (s *partnerService) GetPartner(ctx context.Context, id id.ID) (*PartnerDTO, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	url, _ := s.mediaProvider.GetURL(ctx, p.ResourceID)

	return &PartnerDTO{
		ID:       p.ID,
		Name:     p.Name,
		ImageURL: url,
	}, nil
}

func (s *partnerService) ListPartners(ctx context.Context) ([]*PartnerDTO, error) {
	partners, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var resourceIDs []id.ID
	for _, p := range partners {
		resourceIDs = append(resourceIDs, p.ResourceID)
	}

	urls, _ := s.mediaProvider.GetURLs(ctx, resourceIDs)

	var dtos []*PartnerDTO
	for _, p := range partners {
		dtos = append(dtos, &PartnerDTO{
			ID:       p.ID,
			Name:     p.Name,
			ImageURL: urls[p.ResourceID],
		})
	}

	return dtos, nil
}
