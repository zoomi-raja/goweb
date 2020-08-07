package controllers

import (
	"fmt"
	"net/http"

	"github.com/zoomi-raja/goweb/api/models"
	"github.com/zoomi-raja/goweb/api/responses"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	keys := r.URL.Query()
	//	keys := r.FormValue("limit")
	field := keys.Get("limit")
	fmt.Println(field)
	if posts, err := post.GetAll(); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	} else {
		responses.JSON(w, http.StatusOK, posts)
	}
}

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	// var post models.Post
	fmt.Fprintln(w, "helllo")
}