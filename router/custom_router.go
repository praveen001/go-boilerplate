package router

import (
	"net/http"

	"github.com/praveen001/quest-server/controllers"

	"github.com/gorilla/mux"
)

// CustomRouter wrapped mux router
type CustomRouter struct {
	*mux.Router
	*controllers.AppContext
}

func (cr *CustomRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cr.Router.ServeHTTP(w, r)
}

func (cr *CustomRouter) bindHandler(path string, handler http.HandlerFunc, method string) {
	cr.Router.HandleFunc(path, handler).Methods(method)
}

func (cr *CustomRouter) use(path string, subRouter func(*CustomRouter)) {
	scr := &CustomRouter{cr.PathPrefix(path).Subrouter(), cr.AppContext}
	subRouter(scr)
}

func (cr *CustomRouter) get(path string, handler http.HandlerFunc) {
	cr.bindHandler(path, handler, "GET")
}

func (cr *CustomRouter) post(path string, handler http.HandlerFunc) {
	cr.bindHandler(path, handler, "POST")
}

func (cr *CustomRouter) put(path string, handler http.HandlerFunc) {
	cr.bindHandler(path, handler, "PUT")
}

func (cr *CustomRouter) delete(path string, handler http.HandlerFunc) {
	cr.bindHandler(path, handler, "PUT")
}
