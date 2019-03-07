package router

import (
	"github.com/go-chi/chi"
	"github.com/praveen001/go-boilerplate/handlers"
)

func (cr *CustomRouter) mediaRouter() *chi.Mux {
	media := handlers.NewMediaHandler(cr.appCtx)

	r := chi.NewRouter()
	r.Get("/", media.List)

	return r
}
