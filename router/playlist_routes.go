package router

import (
	"github.com/go-chi/chi"
	"github.com/praveen001/go-boilerplate/handlers"
)

func (cr *CustomRouter) playlistRouter() *chi.Mux {
	playlist := handlers.NewPlaylistHandler(cr.appCtx)

	r := chi.NewRouter()
	r.Post("/date/{date}", playlist.Create)
	r.Get("/date/{date}", playlist.ReadByDate)

	r.Route("/{playlistID}", func(r chi.Router) {
		r.Use(playlist.Preload)
		r.Get("/", playlist.Read)
		r.Put("/", playlist.Update)
		r.Delete("/", playlist.Delete)
	})

	return r
}
