package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/praveen001/go-boilerplate/app"
	"github.com/praveen001/go-boilerplate/handlers"
)

// CustomRouter wrapps chi MUX router, and application context
//
// Allows passing application context to handlers
type CustomRouter struct {
	*chi.Mux
	appCtx *app.Context
}

func (cr *CustomRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cr.Mux.ServeHTTP(w, r)
}

// New initializes the application's router
func New(c *app.Context) http.Handler {
	cr := &CustomRouter{
		chi.NewMux(),
		c,
	}

	base := handlers.NewBaseHandler(c)

	cr.Use(base.DummyAuth, base.CORSHandler, base.LogHandler, base.RecoveryHandler)

	cr.Route("/v2/api", func(r chi.Router) {
		r.Mount("/feeds", cr.feedRouter())
	})

	return cr
}

/*
curl -H "Origin: http://example.com" \
-H "Access-Control-Request-Method: POST" \
-H "Access-Control-Request-Headers: X-Requested-With" \
-X OPTIONS --verbose http://127.0.0.1:5000
*/
