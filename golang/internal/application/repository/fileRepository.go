package repository

import (
	"context"
	"verniy-fm-backend/internal/domain/file"
)

type FileRepository interface {
	ListFiles(context.Context, file.Path) ([]*file.Metadata, error)
}
