package router

import (
	"net/http"

	"github.com/zoomi-raja/goweb/api/controllers"
	"github.com/zoomi-raja/goweb/api/router/routes"

	"github.com/gorilla/mux"
)

// SetupRoutes register all routes to mux
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range routes.Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	r.NotFoundHandler = http.HandlerFunc(controllers.NotFound)
	return r
}

//New init mux router
func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return SetupRoutes(router)
}
