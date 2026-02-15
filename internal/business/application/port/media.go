package port

import (
	"context"

	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type MediaProvider interface {
	GetURL(ctx context.Context, resourceID id.ID) (string, error)
	GetURLs(ctx context.Context, resourceIDs []id.ID) (map[id.ID]string, error)
}
