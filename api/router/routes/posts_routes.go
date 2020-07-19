package routes

import (
	"gotut/api/controllers"
	"net/http"
)

var postsRoutes = []Route{
	Route{
		Uri:     "/posts",
		Method:  http.MethodGet,
		Handler: controllers.GetPosts,
	},
}
