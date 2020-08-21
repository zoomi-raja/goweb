package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetPost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if post, err := post.GetById(id); err != nil {
		statusCode := http.StatusInternalServerError
		if err == sql.ErrNoRows {
			statusCode = http.StatusNotFound
		}
		responses.ERROR(w, statusCode, err)
	} else {
		responses.JSON(w, http.StatusOK, post)
	}
}

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	// var post models.Post
	fmt.Fprintln(w, "helllo")
}
