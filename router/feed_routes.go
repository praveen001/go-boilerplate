package router

import (
	"github.com/go-chi/chi"
	"github.com/praveen001/go-boilerplate/handlers"
)

func (cr *CustomRouter) feedRouter() *chi.Mux {
	feed := handlers.NewFeedHandler(cr.context)

	r := chi.NewRouter()
	r.Post("/", feed.Create)
	r.Get("/", feed.List)
	r.Delete("/", feed.DeleteAll)

	r.Route("/{feedID}", func(r chi.Router) {
		r.Use(feed.Preload)
		r.Get("/", feed.Get)
		r.Put("/", feed.Update)
		r.Delete("/", feed.Delete)
	})

	return r
}
