package router

import (
	"github.com/go-chi/chi"
	"github.com/praveen001/go-boilerplate/handlers"
)

func (cr *CustomRouter) feedRouter() *chi.Mux {
	feed := handlers.NewFeedHandler(cr.context)

	router := chi.NewRouter()

	router.Route("/{feedID}", func(r chi.Router) {
		router.Get("/", feed.Get)
		router.Put("/", feed.Update)
		router.Delete("/", feed.Delete)
	})

	return router
}
