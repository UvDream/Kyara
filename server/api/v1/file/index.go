package file

import "server/service"

type ApiFileGroup struct {
	FilesApi
	ImageApi
}

var (
	imageService = service.ServicesGroupApp.FileServiceGroup.ImageService
)
