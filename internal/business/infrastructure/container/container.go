package container

import (
	"github.com/ming-0x0/yuan/internal/business/application"
	"github.com/ming-0x0/yuan/internal/business/infrastructure/adapter"
	"github.com/ming-0x0/yuan/internal/business/infrastructure/persistence"
	mediaApp "github.com/ming-0x0/yuan/internal/media/application"
)

type Container struct {
	PartnerService application.PartnerService
}

func NewContainer(mediaSvc mediaApp.MediaService) *Container {
	partnerRepo := persistence.NewPartnerRepository()
	mediaProvider := adapter.NewMediaAdapter(mediaSvc)
	partnerSvc := application.NewPartnerService(partnerRepo, mediaProvider)

	return &Container{
		PartnerService: partnerSvc,
	}
}
