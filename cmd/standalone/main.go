package main

import (
	"context"
	"fmt"

	businessContainer "github.com/ming-0x0/yuan/internal/business/infrastructure/container"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
	mediaContainer "github.com/ming-0x0/yuan/internal/media/infrastructure/container"
	"github.com/ming-0x0/yuan/pkg/timezone"
)

func main() {
	timezone.SetTimeZoneICT()

	// 1. Initialize Context Containers
	mediaC := mediaContainer.NewContainer()
	businessC := businessContainer.NewContainer(mediaC.MediaService)

	// 2. Access services from containers
	partnerService := businessC.PartnerService

	// Usage example
	ctx := context.Background()
	p, err := partnerService.GetPartner(ctx, id.MustNew())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else if p != nil {
		fmt.Printf("Partner: %s, Image: %s\n", p.Name, p.ImageURL)
	}

	fmt.Println("System initialized via modular containers")
}
