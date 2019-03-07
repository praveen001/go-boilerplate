package router

import (
	"github.com/go-chi/chi"
)

func (cr *CustomRouter) mediaRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", cr.handler.ListMedias)

	return r
}
