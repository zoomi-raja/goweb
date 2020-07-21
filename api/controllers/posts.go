package controllers

import (
	"gotut/api/models"
	"gotut/api/responses"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if posts, err := post.GetAll(); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	} else {
		responses.JSON(w, http.StatusOK, posts)
	}
}
