package router

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/praveen001/go-boilerplate/controllers"

	"github.com/gorilla/mux"
)

// InitRouter initializes the application's router
func InitRouter(ctx *controllers.AppContext) http.Handler {
	r := &CustomRouter{
		mux.NewRouter(),
		ctx,
	}

	r.use("/users", userRouter)

	return alice.New(ctx.LogHandler, ctx.CORSHandler, ctx.RecoveryHandler).Then(r)
}

/*
curl -H "Origin: http://example.com" \
-H "Access-Control-Request-Method: POST" \
-H "Access-Control-Request-Headers: X-Requested-With" \
-X OPTIONS --verbose http://127.0.0.1:5000
*/
