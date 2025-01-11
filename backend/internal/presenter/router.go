package presenter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	file := &FileHandler{}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Verniy file manager!"))
	})
	r.Route("/api", func(r chi.Router) {
		r.Route("/file", func(r chi.Router) {
			r.Post("/list", file.ListFiles)
		})
	})
	return r
}
