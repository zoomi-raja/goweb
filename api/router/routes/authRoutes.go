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
		Route{
			Uri:     "/logup",
			Method:  http.MethodPost,
			Handler: controllers.CreateUser,
		},
	}
	routes = append(routes, aRoutes...)
}
