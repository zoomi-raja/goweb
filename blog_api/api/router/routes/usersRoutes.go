package routes

import (
	"net/http"

	"github.com/zoomi-raja/goweb/api/middleware"

	"github.com/zoomi-raja/goweb/api/controllers"
)

func init() {
	middlewares := []middleware.Middleware{
		middleware.SetMiddlewareLogger,
		middleware.AuthMiddleware,
	}
	uRoutes := []Route{
		Route{
			Uri:     "/users",
			Method:  http.MethodGet,
			Handler: middleware.Addmiddlewares(controllers.GetUsers, middlewares),
		},
	}
	routes = append(routes, uRoutes...)
}
