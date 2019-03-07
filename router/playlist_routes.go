package router

import (
	"github.com/go-chi/chi"
)

func (cr *CustomRouter) playlistRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", cr.handler.CreatePlaylist)
	r.Get("/date/{date}", cr.handler.GetPlaylistByDate)

	r.Route("/{playlistID}", func(r chi.Router) {
		r.Use(cr.handler.PreloadPlaylist)
		r.Get("/", cr.handler.GetPlaylist)
		r.Put("/", cr.handler.UpdatePlaylist)
		r.Delete("/", cr.handler.DeletePlaylist)
	})

	return r
}
