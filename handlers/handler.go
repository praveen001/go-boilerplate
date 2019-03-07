package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/praveen001/go-boilerplate/app"
	"github.com/praveen001/go-boilerplate/handlers/ctx"
	"github.com/praveen001/go-boilerplate/models"
	"github.com/rs/cors"
)

// Handler .
type Handler struct {
	*app.Context
}

// RecoveryHandler returns 500 status when handler panics.
// Writes error to application log
func (h *Handler) RecoveryHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if e := recover(); e != nil {
				switch t := e.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				h.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// CORSHandler handles cors requests
func (h *Handler) CORSHandler(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedHeaders: []string{"authorization", "content-type"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}).Handler(next)
}

// LogHandler writes access log
func (h *Handler) LogHandler(next http.Handler) http.Handler {
	return middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: h.Logger, NoColor: true})(next)
}

// DummyAuth .
func (h *Handler) DummyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := ctx.SetUser(r.Context(), &models.User{ID: 4})

		next.ServeHTTP(w, r.WithContext(c))
	})
}
