package disk

import (
	"context"
	"log/slog"
	"os"
	"verniy-fm-backend/internal/domain/file"
)

type LocalDisk struct{}

func NewLocalDisk() *LocalDisk {
	return &LocalDisk{}
}

func (*LocalDisk) ListFiles(ctx context.Context, dir file.Path) ([]*file.Metadata, error) {
	files, err := os.ReadDir(dir.String())
	if err != nil {
		return nil, err
	}

	var metadata []*file.Metadata
	for _, f := range files {
		i, err := f.Info()
		if err != nil {
			slog.WarnContext(ctx, "failed to get file info", "error", err.Error())
			continue
		}

		metadata = append(metadata, &file.Metadata{
			Name:    f.Name(),
			Size:    i.Size(),
			IsDir:   i.IsDir(),
			ModTime: i.ModTime(),
		})
	}

	return metadata, nil
}
