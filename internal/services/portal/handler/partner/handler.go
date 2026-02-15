package partner

import (
	"context"

	businessApp "github.com/ming-0x0/yuan/internal/modules/business/application"
	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
	mediaApp "github.com/ming-0x0/yuan/internal/modules/media/application"
)

type PartnerResponse struct {
	ID       id.ID  `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type PartnerHandler struct {
	businessService businessApp.PartnerService
	mediaService    mediaApp.MediaService
}

func NewPartnerHandler(
	businessService businessApp.PartnerService,
	mediaService mediaApp.MediaService,
) *PartnerHandler {
	return &PartnerHandler{
		businessService: businessService,
		mediaService:    mediaService,
	}
}

func (h *PartnerHandler) GetPartner(ctx context.Context, partnerID id.ID) (*PartnerResponse, error) {
	// 1. Get core data from Business Module
	p, err := h.businessService.GetPartner(ctx, partnerID)
	if err != nil {
		return nil, err
	}

	// 2. Get additional data from Media Module
	imageURL := ""
	res, _ := h.mediaService.GetResource(ctx, p.ResourceID)
	if res != nil {
		imageURL = res.URL
	}

	// 3. Combine into a site-specific Response/DTO
	return &PartnerResponse{
		ID:       p.ID,
		Name:     p.Name,
		ImageURL: imageURL,
	}, nil
}
