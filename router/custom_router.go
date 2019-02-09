package router

import (
	"net/http"

	"github.com/praveen001/go-boilerplate/controllers"

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

func (cr *CustomRouter) subRouter(path string, fn func(*CustomRouter)) {
	fn(&CustomRouter{cr.PathPrefix(path).Subrouter(), cr.AppContext})
}

func (cr *CustomRouter) get(path string, handler http.HandlerFunc) {
	cr.bindHandler(path, handler, http.MethodGet)
}

func (cr *CustomRouter) post(path string, handler http.HandlerFunc) {
	cr.bindHandler(path, handler, http.MethodPost)
}

func (cr *CustomRouter) put(path string, handler http.HandlerFunc) {
	cr.bindHandler(path, handler, http.MethodPut)
}

func (cr *CustomRouter) delete(path string, handler http.HandlerFunc) {
	cr.bindHandler(path, handler, http.MethodDelete)
}
