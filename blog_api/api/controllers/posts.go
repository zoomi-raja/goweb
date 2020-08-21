package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/zoomi-raja/goweb/api/models"
	"github.com/zoomi-raja/goweb/api/requests"
	"github.com/zoomi-raja/goweb/api/responses"
)

/*GetPosts returns list of blog posts available in our system*/
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

/*GetPost is handler to get particular post against post id*/
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

/*CreatePosts is handler for post request on create*/
func CreatePosts(w http.ResponseWriter, r *http.Request) {
	post := requests.Post{}
	json.NewDecoder(r.Body).Decode(&post)
	err := post.ValidatePostCreate()
	if err.HasError() {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	values := r.Context().Value(models.AuthCTX("user")).(jwt.MapClaims)
	id, _ := values["user_id"]
	postModel := models.Post{UserID: int64(id.(float64)), Title: post.Title, Intro: post.Intro, Body: post.Body}
	if postID, err := postModel.CreatePost(); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	} else {
		responses.JSON(w, http.StatusOK, struct {
			PostID int64 `json:"postID"`
		}{postID})
	}
}
