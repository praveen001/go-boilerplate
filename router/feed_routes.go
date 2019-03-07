package router

import (
	"github.com/go-chi/chi"
)

func (cr *CustomRouter) feedRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", cr.handler.ListFeeds)

	r.Route("/{feedID}", func(r chi.Router) {
		r.Use(cr.handler.PreloadFeed)
		r.Get("/", cr.handler.GetFeed)

		r.Mount("/playlists", cr.playlistRouter())
		r.Mount("/medias", cr.mediaRouter())
	})

	return r
}
