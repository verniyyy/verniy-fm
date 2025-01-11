package presenter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Verniy file manager!"))
	})

	return r
}
