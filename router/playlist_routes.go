package router

import (
	"github.com/go-chi/chi"
	"github.com/praveen001/go-boilerplate/handlers"
)

func (cr *CustomRouter) playlistRouter() *chi.Mux {
	playlist := handlers.NewPlaylistHandler(cr.context)

	r := chi.NewRouter()
	r.Post("/", playlist.Create)
	r.Get("/date/{date}", playlist.GetByDate)

	r.Route("/{playlistID}", func(r chi.Router) {
		r.Use(playlist.Preload)
		r.Get("/", playlist.Get)
		r.Put("/", playlist.Update)
		r.Delete("/", playlist.Delete)
	})

	return r
}
