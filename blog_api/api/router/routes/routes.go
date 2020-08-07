package routes

import "net/http"

//Route type to hold route parms
type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

var routes []Route

//Load all app urls here
func Load() []Route {
	// routes := append(postsRoutes, usersRoutes...)
	// append other route here to make one slice
	return routes
}
