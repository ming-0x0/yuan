package persistence

import (
	"context"

	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
	"github.com/ming-0x0/yuan/internal/modules/media/domain"
	"github.com/ming-0x0/yuan/internal/modules/media/domain/resource"
)

type resourceRepository struct{}

func NewResourceRepository() domain.ResourceRepository {
	return &resourceRepository{}
}

func (r *resourceRepository) FindByID(ctx context.Context, id id.ID) (resource.Resource, error) {
	return nil, nil // TODO: Implement
}

func (r *resourceRepository) FindByIDs(ctx context.Context, ids []id.ID) ([]resource.Resource, error) {
	return nil, nil // TODO: Implement
}

func (r *resourceRepository) Save(ctx context.Context, res resource.Resource) error {
	return nil // TODO: Implement
}
