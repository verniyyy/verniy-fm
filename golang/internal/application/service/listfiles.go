package service

import (
	"context"
	"verniy-fm-backend/internal/application/repository"
	"verniy-fm-backend/internal/domain/file"
)

type ListFiles struct {
	FileRepo repository.FileRepository
}

func (u ListFiles) Execute(ctx context.Context, input *ListFilesInput) (*ListFilesOutput, error) {
	metadata, err := u.FileRepo.ListFiles(ctx, input.Path)
	if err != nil {
		return nil, err
	}

	return &ListFilesOutput{
		Metadata: metadata,
	}, nil
}

type ListFilesInput struct {
	Path file.Path
}

type ListFilesOutput struct {
	Metadata []*file.Metadata
}
