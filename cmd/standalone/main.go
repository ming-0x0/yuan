package main

import (
	"context"
	"fmt"

	businessContainer "github.com/ming-0x0/yuan/internal/modules/business/infrastructure/container"
	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
	mediaContainer "github.com/ming-0x0/yuan/internal/modules/media/infrastructure/container"
	"github.com/ming-0x0/yuan/pkg/timezone"
)

func main() {
	timezone.SetTimeZoneICT()

	// 1. Initialize Context Containers (Independent)
	mediaC := mediaContainer.NewContainer()
	businessC := businessContainer.NewContainer()

	// 2. Direct Orchestration at Service Layer
	ctx := context.Background()
	partnerID := id.MustNew()

	// Step 1: Call Business Module
	p, err := businessC.PartnerService.GetPartner(ctx, partnerID)
	if err != nil {
		fmt.Printf("Error fetching partner: %v\n", err)
		return
	}

	if p != nil {
		// Step 2: Call Media Module directly to get URL (No Port/Adapter needed)
		res, _ := mediaC.MediaService.GetResource(ctx, p.ResourceID)

		imageURL := ""
		if res != nil {
			imageURL = res.URL
		}

		fmt.Printf("Orchestrated Data -> Partner: %s, Image: %s\n", p.Name, imageURL)
	}

	fmt.Println("System initialized with direct orchestration")
}
