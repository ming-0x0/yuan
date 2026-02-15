package application

import (
	"context"

	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
	"github.com/ming-0x0/yuan/internal/modules/media/domain"
	"github.com/ming-0x0/yuan/internal/modules/media/domain/resource"
)

type MediaService interface {
	GetResource(ctx context.Context, id id.ID) (resource.Resource, error)
	GetResources(ctx context.Context, ids []id.ID) ([]resource.Resource, error)
}

type mediaService struct {
	resRepo domain.ResourceRepository
}

func NewMediaService(resRepo domain.ResourceRepository) MediaService {
	return &mediaService{
		resRepo: resRepo,
	}
}

func (s *mediaService) GetResource(ctx context.Context, id id.ID) (resource.Resource, error) {
	return s.resRepo.FindByID(ctx, id)
}

func (s *mediaService) GetResources(ctx context.Context, ids []id.ID) ([]resource.Resource, error) {
	return s.resRepo.FindByIDs(ctx, ids)
}
