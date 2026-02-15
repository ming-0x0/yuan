package container

import (
	"github.com/ming-0x0/yuan/internal/modules/business/application"
	"github.com/ming-0x0/yuan/internal/modules/business/infrastructure/persistence"
)

type Container struct {
	PartnerService application.PartnerService
}

func NewContainer() *Container {
	partnerRepo := persistence.NewPartnerRepository()
	partnerSvc := application.NewPartnerService(partnerRepo)

	return &Container{
		PartnerService: partnerSvc,
	}
}
