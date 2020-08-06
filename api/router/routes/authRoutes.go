package routes

import (
	"net/http"

	"github.com/zoomi-raja/goweb/api/controllers"
)

func init() {
	var aRoutes = []Route{
		Route{
			Uri:     "/login",
			Method:  http.MethodPost,
			Handler: controllers.Login,
		},
	}
	routes = append(routes, aRoutes...)
}
