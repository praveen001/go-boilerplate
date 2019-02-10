package router

import (
	"github.com/go-chi/chi"
)

func (cr *CustomRouter) userRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/register", cr.RegisterUser)
	router.Post("/register", cr.RegisterUser)

	return router
}
