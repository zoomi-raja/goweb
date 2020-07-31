package controllers

import (
	"fmt"
	"gotut/api/models"
	"gotut/api/responses"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	keys, ok := r.URL.Query()["limit"]
	if ok {
		fmt.Println(keys[0])
	}
	if posts, err := post.GetAll(); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	} else {
		responses.JSON(w, http.StatusOK, posts)
	}
}

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	// var post models.Post
}
