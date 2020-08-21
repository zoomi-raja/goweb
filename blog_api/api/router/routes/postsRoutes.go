package routes

import (
	"net/http"

	"github.com/zoomi-raja/goweb/api/controllers"
	"github.com/zoomi-raja/goweb/api/middleware"
)

func init() {
	middlewares := []middleware.Middleware{
		middleware.AuthMiddleware,
	}

	var pRoutes = []Route{
		Route{
			Uri:     "/posts",
			Method:  http.MethodGet,
			Handler: controllers.GetPosts,
		},
		Route{
			Uri:     "/posts/{id}",
			Method:  http.MethodGet,
			Handler: controllers.GetPost,
		},
		Route{
			Uri:     "/posts",
			Method:  http.MethodPost,
			Handler: middleware.Addmiddlewares(controllers.CreatePosts, middlewares),
		},
	}
	routes = append(routes, pRoutes...)
}
