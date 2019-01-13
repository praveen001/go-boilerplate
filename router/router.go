package router

import (
	"net/http"

	"github.com/praveen001/quest-server/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// InitRouter initializes the application's router
func InitRouter(ctx *controllers.AppContext) http.Handler {
	r := &CustomRouter{
		mux.NewRouter(),
		ctx,
	}

	r.use("/users", userRouter)

	return cors.New(cors.Options{
		AllowedHeaders: []string{"authorization", "content-type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(r)
}
