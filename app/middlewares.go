package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

// RecoveryHandler returns 500 status when handler panics.
// Writes error to application log
func (c *Context) RecoveryHandler(h http.Handler) http.Handler {
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
				c.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		h.ServeHTTP(w, r)
	})
}

// CORSHandler handles cors requests
func (c *Context) CORSHandler(h http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedHeaders: []string{"authorization", "content-type"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}).Handler(h)
}

// LogHandler writes access log
func (c *Context) LogHandler(h http.Handler) http.Handler {
	return middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: c.Logger, NoColor: true})(h)
}

// DummyAuth .
func (c *Context) DummyAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Added userID")
		ctx := context.WithValue(r.Context(), "userID", 2)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
