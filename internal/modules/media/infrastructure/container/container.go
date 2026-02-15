package container

import (
	"github.com/ming-0x0/yuan/internal/modules/media/application"
	"github.com/ming-0x0/yuan/internal/modules/media/infrastructure/persistence"
)

type Container struct {
	MediaService application.MediaService
}

func NewContainer() *Container {
	resRepo := persistence.NewResourceRepository()
	mediaSvc := application.NewMediaService(resRepo)

	return &Container{
		MediaService: mediaSvc,
	}
}
