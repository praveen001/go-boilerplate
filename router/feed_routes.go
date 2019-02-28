package router

import (
	"github.com/go-chi/chi"
	"github.com/praveen001/go-boilerplate/handlers"
)

func (cr *CustomRouter) feedRouter() *chi.Mux {
	feed := handlers.NewFeedHandler(cr.context)

	r := chi.NewRouter()
	r.Get("/", feed.List)

	r.Route("/{feedID}", func(r chi.Router) {
		r.Use(feed.Preload)
		r.Get("/", feed.Get)
	})

	return r
}
