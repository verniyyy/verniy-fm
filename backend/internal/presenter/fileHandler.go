package presenter

import (
	"net/http"
	"verniy-fm-backend/internal/application/service"
	"verniy-fm-backend/internal/domain/file"
	"verniy-fm-backend/internal/infrastructure/injector"
)

type FileHandler struct{}

func (h *FileHandler) ListFiles(w http.ResponseWriter, r *http.Request) {
	input, err := NewListFilesInput(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	service := injector.NewListFilesService()
	output, err := service.Execute(r.Context(), input)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	OK(w, output)
}

func NewListFilesInput(r *http.Request) (*service.ListFilesInput, error) {
	var valErrs []error

	path, err := file.ParsePath(r.URL.Query().Get("path"))
	if err != nil {
		valErrs = append(valErrs, err)
	}

	if len(valErrs) > 0 {
		return nil, NewClientError("invalid input", valErrs...)
	}

	return &service.ListFilesInput{
		Path: path,
	}, nil
}
