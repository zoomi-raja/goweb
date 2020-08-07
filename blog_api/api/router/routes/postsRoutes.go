package routes

import (
	"net/http"

	"github.com/zoomi-raja/goweb/api/controllers"
)

func init() {
	var pRoutes = []Route{
		Route{
			Uri:     "/posts",
			Method:  http.MethodGet,
			Handler: controllers.GetPosts,
		},
		Route{
			Uri:     "/posts",
			Method:  http.MethodPost,
			Handler: controllers.CreatePosts,
		},
	}
	routes = append(routes, pRoutes...)
}
