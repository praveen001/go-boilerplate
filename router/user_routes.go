package router

import (
	"github.com/go-chi/chi"
	"github.com/praveen001/go-boilerplate/handlers"
)

func (cr *CustomRouter) userRouter() *chi.Mux {
	user := handlers.NewUserHandler(cr.context)

	r := chi.NewRouter()
	r.Post("/", user.Create)
	r.Get("/", user.List)
	r.Delete("/", user.DeleteAll)

	r.Group(func(r chi.Router) {
		r.Route("/{userID}", func(r chi.Router) {
			r.Use(user.Preload)
			r.Get("/", user.Get)
			r.Put("/", user.Update)
			r.Delete("/", user.Delete)
		})
	})

	return r
}
