package file

import "server/service"

type ApiFileGroup struct {
	FilesApi
}

var (
	fileService = service.ServicesGroupApp.FileServiceGroup.FilesService
)
