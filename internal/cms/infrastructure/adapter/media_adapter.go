package adapter

import (
	"context"

	"github.com/ming-0x0/yuan/internal/cms/application/port"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/media/application"
)

type mediaAdapter struct {
	mediaService application.MediaService
}

func NewMediaAdapter(mediaService application.MediaService) port.MediaProvider {
	return &mediaAdapter{
		mediaService: mediaService,
	}
}

func (a *mediaAdapter) GetURL(ctx context.Context, resourceID id.ID) (string, error) {
	res, err := a.mediaService.GetResource(ctx, resourceID)
	if err != nil {
		return "", err
	}
	return res.URL, nil
}

func (a *mediaAdapter) GetURLs(ctx context.Context, resourceIDs []id.ID) (map[id.ID]string, error) {
	resources, err := a.mediaService.GetResources(ctx, resourceIDs)
	if err != nil {
		return nil, err
	}

	urls := make(map[id.ID]string)
	for _, res := range resources {
		urls[res.ID] = res.URL
	}
	return urls, nil
}
