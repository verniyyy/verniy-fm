package injector

import (
	"verniy-fm-backend/internal/application/service"
	"verniy-fm-backend/internal/infrastructure/disk"
)

func NewListFilesService() *service.ListFiles {
	return &service.ListFiles{
		FileRepo: disk.NewLocalDisk(),
	}
}
