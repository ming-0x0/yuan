package domain

import (
	"context"

	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/media/domain/resource"
)

type ResourceRepository interface {
	FindByID(ctx context.Context, id id.ID) (resource.Resource, error)
	FindByIDs(ctx context.Context, ids []id.ID) ([]resource.Resource, error)
	Save(ctx context.Context, res resource.Resource) error
}
