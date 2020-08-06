package routes

import (
	"net/http"

	"github.com/zoomi-raja/goweb/api/controllers"
)

func init() {
	uRoutes := []Route{
		Route{
			Uri:     "/users",
			Method:  http.MethodGet,
			Handler: controllers.GetUsers,
		},
		Route{
			Uri:     "/users",
			Method:  http.MethodPost,
			Handler: controllers.CreateUser,
		},
	}
	routes = append(routes, uRoutes...)
}
