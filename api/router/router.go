package router

import (
	"gotut/api/middleware"
	"gotut/api/router/routes"

	"github.com/gorilla/mux"
)

// SetupRoutes register all routes to mux
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range routes.Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}
func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range routes.Load() {
		r.HandleFunc(route.Uri,
			middleware.SetMiddlewareLogger(
				middleware.SetMiddlewareJSON(route.Handler),
			),
		).Methods(route.Method)
	}
	return r
}

//New init mux router
func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return SetupRoutes(router)
}
