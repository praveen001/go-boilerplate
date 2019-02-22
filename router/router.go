package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/praveen001/go-boilerplate/app"
)

// CustomRouter wrapps chi MUX router, and application context
//
// Allows passing application context to handlers
type CustomRouter struct {
	*chi.Mux
	context *app.Context
}

func (cr *CustomRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cr.Mux.ServeHTTP(w, r)
}

// New initializes the application's router
func New(ctx *app.Context) http.Handler {
	cr := &CustomRouter{
		chi.NewMux(),
		ctx,
	}

	cr.Use(ctx.CORSHandler, ctx.LogHandler, ctx.RecoveryHandler)

	cr.Route("/v1/api", func(r chi.Router) {
		r.Mount("/users", cr.userRouter())
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
